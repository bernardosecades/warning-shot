# Dependencies: go mod

Go Mod:

https://github.com/golang/go/wiki/Modules#modules

Generate go.mod:

```go mod init``` 

Downloads dependencies:

```go build```

Generate vendor:

```env GO111MODULE=on go mod vendor```

# FORMAT CODE

```go fmt main.go```

# BUILD

```go build```

# RUN

```PORT=7000 ./warning-shot```

or

```source .env```
```./warning-shot```

# CODE REVIEWS GOLANG
``
https://github.com/golang/go/wiki/CodeReviewComments

# Set the build id using git´s SHA
``
You can get the short version of the sha1 of your latest commit by running the following git command from your repo:

```git rev-parse --short HEAD```
```go build -ldflags "-X main.Build=a1064bc" main.go```


# TEST

```go test <folder>```

```go test ./config```

Text format:

```go test ./config -coverprofile fmt```

Html format:

```go test ./config -coverprofile coverage/cover.out```
```go tool cover -html=coverage/cover.out -o coverage/cover.html```