# Go Hello

Just a Hello World service for tests

## How run from command line

If you have golang on your machine

```
$ go get -v github/efrenfuentes/go_hello

$ cd $GOPATH/src/github/efrenfuentes/go_hello

$ go build

$ ./go_hello
2018/02/18 03:13:50 Server starting on port 8080
```

Go to your browser on http://localhost:8080

## How run it with docker

If you have docker on your machine

```
$ go get -v github/efrenfuentes/go_hello

$ cd $GOPATH/src/github/efrenfuentes/go_hello

$ ./build.sh
Sending build context to Docker daemon  6.144kB
Step 1/8 : FROM golang:1.9-alpine
 ---> bbab7aea1231
Step 2/8 : RUN apk add --no-cache git mercurial
 ---> Using cache
 ---> 595b0b7df48b
Step 3/8 : WORKDIR /go/src/github/efrenfuentes/go_hello
 ---> Using cache
 ---> 8f37a7e0005d
Step 4/8 : COPY . .
 ---> 72f1d861ca58
Step 5/8 : RUN go get
 ---> Running in 692077e3703b
Removing intermediate container 692077e3703b
 ---> 1bdabf5ca66c
Step 6/8 : RUN go build
 ---> Running in ef392a34ed4f
Removing intermediate container ef392a34ed4f
 ---> 04b71513f80b
Step 7/8 : EXPOSE 8080
 ---> Running in 1d3c448893f1
Removing intermediate container 1d3c448893f1
 ---> a4c86b59c6d1
Step 8/8 : ENTRYPOINT ["./go_hello"]
 ---> Running in 3d5361eca1ee
Removing intermediate container 3d5361eca1ee
 ---> dbe6f5063e47
Successfully built dbe6f5063e47
Successfully tagged go_hello:latest

$ ./run.sh
2018/02/18 03:13:50 Server starting on port 8080
```

Go to your browser on http://localhost:8080
