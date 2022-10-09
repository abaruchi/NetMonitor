package monitor

import (
    "io"
    "net/http"
    "sync"
    "time"
)

func SpeedCalculator(urls ContinentHosts) map[string]float64 {

    var mutex sync.Mutex
    var wg sync.WaitGroup

    mapHostsList := ContinentHostsToMap(urls)
    mapHostTests := make(map[string]float64)

    for k, v := range mapHostsList {
        wg.Add(1)
        go func(cont string, hosts []string ) {
            mutex.Lock()
            mapHostTests[cont] = runTest(hosts, getURL)
            mutex.Unlock()
            wg.Done()
        }(k, v)
    }
    wg.Wait()
    return mapHostTests
}

func getURL(u string) float64 {

    start := time.Now()
    res, err := http.Get(u)
    if err != nil {
        return 0
    }
    elapsed := time.Now().Sub(start)

    if res.StatusCode != http.StatusOK {
        return 0
    }

    b, _ := io.ReadAll(res.Body)
    if b == nil {
        return 0
    }
    s := bytesToMb(b) / elapsed.Seconds()
    return s
}

func bytesToMb(b []byte) float64 {
    return float64(len(b)) / 1024 / 1024
}
