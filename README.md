# echo-key

Small and simple echo server that will respond with the value of key when it is requested. This can be used to run a server in a docker build to keep secrets out of the build layer. If you want to use this with multiple keys, just run multiple instances and bind to different ports.
 
The Docker image runs in an Alpine Linux instance and is only 11.5 MB in size when built.

##Usage

Run Docker
```
docker run -d -p 9000:8080 -e key="hello world" echo-key
```

curl (use your ip address)
```
curl 192.168.99.100:9000
```

Build standalone for Alpine

```
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang sh -c 'CGO_ENABLED=0 GOOS=linux go build -a --installsuffix cgo --ldflags="-s -w" -o echo-key'
```