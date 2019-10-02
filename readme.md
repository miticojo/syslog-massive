# SYSLOG UDP/TCP TEST   

## BUILD
With make, golang and docker installed, run:
```bash
 make
```

## RUN

### SERVER
default values for are:
* PROTO=udp
* PORT=2000

to run it with no values specified:
```
  docker run -p 2000:2000/udp -ti miticojo/syslog-massive-server:latest
```
or you can specify port and protocol:
```
  docker run -p 3000:3000 -ti -e PORT=3000 -e PROTO=tcp miticojo/syslog-massive-server:latest
```
Server daemon always listen on 0.0.0.0.

### CLIENT
default values for are:
* PROTO=udp
* PORT=2000

to run it with no values specified:
```
docker run --rm miticojo/syslog-massive-server:latest ./syslog-massive-client 10 udp 192.168.1.10:2000
```
