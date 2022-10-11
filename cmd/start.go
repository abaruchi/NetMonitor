/*
Copyright © 2022 Artur Baruchi <abaruchi@abaruchi.dev>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"github.com/abaruchi/NetMonitor/pkg/monitor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var (
	varHostList     string
	promEndpoint    string
	promHttpPort    string
	refreshInterval time.Duration
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start NetMon tool.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")

		// Start prometheus handler
		err := prometheus.Register(monitor.Latency)
		if err != nil {
			fmt.Printf("Error registering latency")
		}
		err = prometheus.Register(monitor.Speed)
		if err != nil {
			fmt.Printf("Error registering speed")
		}

		go func() {
			fmt.Printf("Starting Prometheus: Port %s, Endpoint %s\n", promHttpPort, promEndpoint)
			http.Handle(promEndpoint, promhttp.Handler())
			err := http.ListenAndServe(":"+promHttpPort, nil)
			if err != nil {
				fmt.Printf("Error to start prometheus metrics: %s\n", err)
				return
			}
		}()
		for range time.Tick(45 * time.Second) {
			runMonitor()
		}

	},
}

func runMonitor() {

	tempContHosts := monitor.ContinentHosts{
		NorthAmerica: []string{"https://www.unam.mx/", "https://www.columbia.edu/", "https://www.ucalgary.ca/", "https://www.fiu.edu/"},
		SouthAmerica: []string{"https://www.poli.usp.br/", "https://www.pucsp.br/home", "https://www.unicen.edu.ar/", "http://www.unla.edu.ar/"},
		Oceania:      []string{"https://www.nsw.gov.au/", "https://www.unimelb.edu.au/", "https://www.telstra.com.au/", "https://www.sydney.edu.au/"},
		Asia:         []string{"https://www.hanyang.ac.kr/", "http://hust.edu.cn/index.htm", "https://www.sustech.edu.cn/", "https://www.cau.ac.kr/index.do", "https://web.ncku.edu.tw/"},
		Europe:       []string{"https://www.jacobs-university.de/", "https://www.ulisboa.pt/", "https://www.uminho.pt/PT", "http://www.tu-dresden.de/"},
		Africa:       []string{"https://www.uct.ac.za/", "https://www.up.ac.za/", "https://www.uonbi.ac.ke/", "https://www.unam.edu.na/"},
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		t := monitor.SpeedCalculator(tempContHosts)
		wg.Done()
		updateSpeedMetrics(t)
	}()

	wg.Add(1)
	go func() {
		l := monitor.LantencyAvg(tempContHosts)
		wg.Done()
		updateLatencyMetrics(l)
	}()
	wg.Wait()
	return
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVar(&varHostList, "host-list", "", "Optional list of hosts to use for testing.")
	startCmd.Flags().StringVar(&promEndpoint, "prom-endpoint", "/metrics", "Endpoint to be scrapped by Prometheus.")
	startCmd.Flags().StringVar(&promHttpPort, "prom-port", "9001", "HTTP Port to be exported and scrapped by Prometheus.")
	startCmd.Flags().DurationVar(&refreshInterval, "fresh-interval", 3 * time.Minute, "Fresh rate to gather metrics.")
}

func updateSpeedMetrics(m map[string]float64) {
	for cont, val := range m {
		monitor.Speed.WithLabelValues(cont).Set(val)
	}
}

func updateLatencyMetrics(m map[string]float64) {
	for cont, val := range m {
		monitor.Latency.WithLabelValues(cont).Set(val)
	}
}