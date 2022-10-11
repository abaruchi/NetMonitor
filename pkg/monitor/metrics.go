package monitor

import (
    "github.com/prometheus/client_golang/prometheus"
)

var (
    Latency = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "icmp_latency",
            Help: "Average Latency, in ms, to each Continent.",
        },
        []string{"continent"},
    )
    Speed = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "download_speed",
            Help: "Average Download speed, in mb/s, from each Continent.",
        },
        []string{"continent"},
    )
)
