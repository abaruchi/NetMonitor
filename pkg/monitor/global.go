package monitor

const (
    numberOfContinents = 5
    threadsNumber      = 2
)


// tester is a type that any tester function (i.e. latency or speed) should comply.
type tester func(h string) float64

// ContinentHosts represents the Input used to process and evaluate the latency and speed for each continent.
type ContinentHosts struct {
    America []string
    Oceania []string
    Asia    []string
    Europe  []string
    Africa  []string
}

// HostsAvg this data struct will store the average value for each value computed.
type HostsAvg struct {
    LatencyAvg       float64
    DownloadSpeedAvg float64
}

// ContinentalEvaluation this is the final goal. After measuring each host and taking the average, we want to output this
// struct.
type ContinentalEvaluation struct {
    America HostsAvg
    Oceania HostsAvg
    Asia    HostsAvg
    Europe  HostsAvg
    Africa  HostsAvg
}

func calculateAvg(values []float64) float64 {
    s := 0.0
    l := 0

    // If the value is 0 we don't want to compute it in our Average.
    for _, v := range values {
        if v != 0 {
            l += 1
            s += v
        }
    }
    return s/float64(l)
}

func runTest(hosts []string, tf tester) float64 {
    testResults := make([]float64, len(hosts))
    for i, h := range hosts {
        testResults[i] = tf(h)
    }
    return calculateAvg(testResults)
}

func ContinentHostsToMap(ch ContinentHosts) map[string][]string {

    chMap := make(map[string][]string)
    if len(ch.Asia) > 0 {
        chMap["Asia"] = ch.Asia
    }

    if len(ch.America) > 0 {
        chMap["America"] = ch.America
    }

    if len(ch.Africa) > 0 {
        chMap["Africa"] = ch.Africa
    }

    if len(ch.Oceania) > 0 {
        chMap["Oceania"] = ch.Oceania
    }

    if len(ch.Europe) > 0 {
        chMap["Europe"] = ch.Europe
    }

    return chMap
}
