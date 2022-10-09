package monitor

import (
    "github.com/gandaldf/ping"
    "net/url"
    "sync"
    "time"
)
const (
    pingCount = 5
    pingTimeOut = 3 * time.Second
)

func LantencyAvg(urls ContinentHosts) map[string]float64 {

    var mutex sync.Mutex
    var wg sync.WaitGroup

    mapHostsList := ContinentHostsToMap(urls)

    mapHostTests := make(map[string]float64)
    for k, v := range mapHostsList {
        wg.Add(1)
        go func(cont string, hosts []string) {
            mutex.Lock()
            mapHostTests[cont] = runTest(hosts, pingHost)
            mutex.Unlock()
            wg.Done()
        }(k, v)
    }
    wg.Wait()
    return mapHostTests
}

func pingHost(host string) float64 {

    hostToPing, _ := getHostFromURL(host)

    pinger, err := ping.NewPinger(hostToPing)
    if err != nil {
        return 0
    }
    pinger.Count = pingCount
    pinger.Timeout = pingTimeOut

    err = pinger.Run()
    if err != nil {
        return 0
    }
    stats := pinger.Statistics()
    return float64(stats.AvgRtt.Milliseconds())
}

func getHostFromURL(u string) (string, error) {
    h, err := url.Parse(u)
    if err != nil {
        return "", err
    }
    return h.Host, nil
}
