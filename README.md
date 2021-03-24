This is simple tcp port scan tools, only support ip v4

* complie
1. clone code
2. cd PortScan
3. bash build.sh

* example 
scan 192.168.0.1 to 192.168.255.255 ,80 or 443 is opend or closed

* run 

./portscan  -a_start=192 -a_end=192 -b_start=168 -b_end=168 -c_start=0 -c_end=255 -d_start=1 -d_end=255 -ports=80,443 -timeout=1000ms

* run result logger
```
version is v0.0.1

astart=192, aend=192, bstart=168, bend=168, cstart=0, cend=255,  dstart=1, dend=255,  timeout=1000ms
count=0, address=192.168.0.2:80, opend=false, startTime=1616596193688869000, stopTime=1616596194695025000, timeout=1s, times=1006, totalTimes=1006, err=<nil>
```

* debug or source running

go run Main.go -a_start=192 -a_end=192 -b_start=168 -b_end=168 -c_start=0 -c_end=255 -d_start=1 -d_end=255 -ports=80,443 -timeout=1000ms

* worker process

address and port sets ---> go routing pool ---> send tcp handshake ---> handshake success is opened port otherwise is closed. 