* notes
this is simple tcp port scan tools, only support ip v4

* complie
1. clone code
2. cd PortScan
3. bash build.sh

* example 
scan 192.168.0.1 to 192.168.255.255 ,80 or 443 is opend or closed

* run at linux 
./port_scan_for_linux  -a_start=192 -a_end=192 -b_start=168 -b_end=168 -c_start=0 -c_end=255 -d_start=1 -d_end=255 -ports=80,443 -timeout=1000ms

* run at windows
./port_scan_for_windows -a_start=192 -a_end=192 -b_start=168 -b_end=168 -c_start=0 -c_end=255 -d_start=1 -d_end=255 -ports=80,443 -timeout=1000ms

* run at mac
./port_scan_for_mac  -a_start=192 -a_end=192 -b_start=168 -b_end=168 -c_start=0 -c_end=255 -d_start=1 -d_end=255 -ports=80,443 -timeout=1000ms

* debug or source running
go run Main.go -a_start=192 -a_end=192 -b_start=168 -b_end=168 -c_start=0 -c_end=255 -d_start=1 -d_end=255 -ports=80,443 -timeout=1000ms

* worker process

address and port sets ---> go routing pool ---> send tcp handshake ---> handshake success is opened port otherwise is closed. 