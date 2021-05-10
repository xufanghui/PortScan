This is simple tcp port scan tools, only support ip v4

* complie
1. clone code
2. cd PortScan
3. bash build.sh

* example 
scan 192.168.0.1 to 192.168.255.255 ,80 or 443 is opend or closed

* run 

./portscan  -start=192.168.1.1 -end=192.168.1.255 -ports=80,443 -timeout=1000ms

* run result logger
```
version is v0.0.1

start=192.168.1.1, end=192.168.1.255, timeout=1000ms
count=0, address=192.168.0.2:80, opend=false, startTime=1616596193688869000, stopTime=1616596194695025000, timeout=1s, times=1006, totalTimes=1006, err=<nil>
```

* debug or source running

go run Main.go -start=192.168.1.1 -end=192.168.1.255 -ports=80,443 -timeout=1000ms

* worker process

address and port sets ---> go routing pool ---> send tcp handshake ---> handshake success is opened port otherwise is closed. 

#### Don't do evil

* MIT LISENSE
```
MIT License

Copyright (c) 2021 fanghui

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```