# infinity-http-stream

Server with infinity response stream in golang. Can simulate traffic between device running on linux.

## Usage
1. compile for device
```bash
// for your machine
go build
// compile for airFiber
env GOOS=linux GOARM=5 GOARCH=arm go build
// compile for edgeRouter
env GOOS=linux GOMIPS=softfloat GOARCH=mips go build
```
2. copy and run on server device
```bash
scp infinity-http-stream user@ip-address-server:/tmp/infinity-http-stream
ssh user@ip-address-server '/tmp/infinity-http-stream &'
```
3. execute on client device
```bash
ssh uesr@ip-address-client 'curl ip-address-server:3009 > /dev/null &'
```
4. stop speed traffic
```
restart client and server device
```