package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const VERSION = "v0.0.1"

type ScanResult struct {
	address   string `ip address，only support IP v4`
	opened bool `is opened , true opened; false closed`
	incomeQTime int64 `incoming Goroute queue start time, unit is nanosecond`
	socketStartTime int64 `socket start time， unit is nanosecond`
	socketStopTime int64 `socket stopped time， unit is nanosecond`
	err error
	timeout time.Duration
	count int64

}

// go route count
var  routeCount int  = 2000
// task count
var  taskCount int  = 2000
// finished task count
var finishedTaskCount int = 0
// timeout millsecond
var timeout time.Duration = 200 * time.Millisecond
// local address ,example: 0.0.0.0:0
var localAddress *string

func Task(result *ScanResult ) {
	// timeout
	finalTimeout := timeout
	result.socketStartTime = time.Now().UnixNano()

	var conn net.Conn
	var err error

	if localAddress != nil {
		laddr,err:= net.ResolveTCPAddr("tcp",*localAddress)
		if err !=nil {

		}
		d := net.Dialer{LocalAddr: laddr, Timeout: finalTimeout}
		conn, err = d.Dial("tcp", result.address)
	}else{
		conn, err = net.DialTimeout("tcp", result.address, finalTimeout )
	}
	
	var tcpConn *net.TCPConn
	ok := false
	if conn != nil {
		tcpConn, ok = conn.(*net.TCPConn)
	}

	result.opened = ok
	result.socketStopTime = time.Now().UnixNano()
    result.err = err
    result.timeout = finalTimeout

	if tcpConn != nil {
		defer tcpConn.Close()
	}

	return
}

type OneFinishCallBack func(result *ScanResult)

func callBackForConsolePrintln(result *ScanResult){
	fmt.Printf("count=%d, address=%s, opend=%v, startTime=%d, stopTime=%d, timeout=%v, times=%v, totalTimes=%v, err=%v\n", result.count, result.address, result.opened, result.socketStartTime, result.socketStopTime, result.timeout, (result.socketStopTime-result.socketStartTime)/1e6, (result.socketStopTime-result.incomeQTime)/1e6, result.err )
}

func Workers(task func(result *ScanResult), callBack OneFinishCallBack, allFinishedCallBack func() ) chan ScanResult {
	input := make(chan ScanResult)
	ack := make(chan bool)
	for i := 0; i < routeCount; i++ {
		go func() {
			for {
				taskInput, ok := <-input
				if ok {
					task(&taskInput)
					ack <- true
					callBack(&taskInput)
				}
			}
		}()
	}

	go func() {
		for {
			<-ack
			finishedTaskCount = finishedTaskCount+1
			if finishedTaskCount >= taskCount {
				break
			}
		}

		allFinishedCallBack()
	}()
	return input
}

func StartProfile(){
	// brew install graphviz
	// go tool pprof -http=:8080  'http://127.0.0.1:9999/debug/pprof/profile?seconds=60'
	http.ListenAndServe(":9999", nil)
}


func main(){

	astart := flag.Int("a_start", 172, "ip v4 as : a.b.c.d, start a range, as: 0~255")
	aend := flag.Int("a_end",172 , "ip v4 as : a.b.c.d, end a range, as: 0~255")
	bstart := flag.Int("b_start", 16, "ip v4 as : a.b.c.d, start b range, as: 0~255")
	bend := flag.Int("b_end", 17, "ip v4 as : a.b.c.d, end b range, as: 0~255")
	cstart := flag.Int("c_start", 1, "ip v4 as : a.b.c.d, start c range, as: 0~255")
	cend := flag.Int("c_end", 255, "ip v4 as : a.b.c.d, end c range, as: 0~255")
	dstart := flag.Int("d_start", 1, "ip v4 as : a.b.c.d, start d range, as: 0~255")
	dend := flag.Int("d_end", 255, "ip v4 as : a.b.c.d, end d range, as: 0~255")
	input_timeout := flag.String("timeout", "200ms", "timeout millseconds")
	ports1 := flag.String("ports", "80,443", "ports,default 80,443")
	pcount := flag.Int("pcount", 2000, "go route total count")
	laddr := flag.String("laddr", "0.0.0.0:0", "local ip address and port. as: 0.0.0.0:0")

	fmt.Printf("version is %s\n\n", VERSION)
	go  StartProfile()

	flag.Parse()

	if  input_timeout !=nil {
		timeout ,_ = time.ParseDuration(*input_timeout)
	}

	if laddr !=nil {
		localAddress = laddr
	}

	routeCount = int(*pcount)

	fmt.Printf("astart=%d, aend=%d, bstart=%d, bend=%d, cstart=%d, cend=%d,  dstart=%d, dend=%d,  timeout=%s\n",*astart,*aend,*bstart,*bend,*cstart,*cend,*dstart,*dend,*input_timeout);

	exit := make(chan bool)

	workers := Workers(func(result *ScanResult) {
		Task(result)
	}, callBackForConsolePrintln, func() {
		exit <- true
	})



	ports := strings.Split(*ports1,",")

	arange := *aend-*astart +1

	brange := *bend-*bstart +1

	crange := *cend-*cstart +1

	drange := *dend-*dstart +1


	taskCount = arange*brange*crange*drange*len(ports)
    var count int64 = 0
	for a := *astart; a <= *aend; a++{
		for b:= *bstart; b<= *bend; b++{
			for c:= *cstart; c<= *cend; c++{
				for d:= *dstart; d<= *dend; d++{
					for i := range ports{

						address := strconv.Itoa(a)+"."+strconv.Itoa(b)+"."+strconv.Itoa(c)+"."+strconv.Itoa(d)+":"+ports[i]

						scanResult := ScanResult{address,false,time.Now().UnixNano(),0,0, nil, 0, count }
						count = count + 1
						workers <- scanResult
					}

				}
			}
		}
	}

	close(workers)

	<-exit
}