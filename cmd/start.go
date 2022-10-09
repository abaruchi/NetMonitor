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
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var (
	varHostList string
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start NetMon tool.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")

		for range time.Tick(3 * time.Minute) {
			runMonitor()
		}

	},
}

func runMonitor() {

	fmt.Println("Starting Task!\n")
	tempContHosts := monitor.ContinentHosts{
		America: []string{"https://www.unam.mx/", "https://web.mit.edu/", "https://www.ucalgary.ca/", "https://www.lapresse.ca/", "https://www5.usp.br/"},
		Oceania: []string{"https://www.nsw.gov.au/", "https://www.unimelb.edu.au/", "https://www.telstra.com.au/", "https://www.sydney.edu.au/"},
		Asia:    []string{"https://www.hanyang.ac.kr/", "http://hust.edu.cn/index.htm", "https://www.sustech.edu.cn/", "https://www.cau.ac.kr/index.do", "https://web.ncku.edu.tw/"},
		Europe:  []string{"https://www.jacobs-university.de/", "https://www.ulisboa.pt/", "https://www.uminho.pt/PT", "http://www.tu-dresden.de/"},
		Africa:  []string{"https://www.uct.ac.za/", "https://www.up.ac.za/", "https://www.uonbi.ac.ke/", "https://www.unam.edu.na/"},
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		t := monitor.SpeedCalculator(tempContHosts)
		fmt.Printf("SpeedDownload: %v\n", t)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		l := monitor.LantencyAvg(tempContHosts)
		fmt.Printf("Latency: %v\n", l)
		wg.Done()
	}()
	wg.Wait()
	return
}


func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVar(&varHostList, "host-list", "", "Optional list of hosts to use for testing.")
}
