# Install Go

- install
```shell
brew install go
```

- check version
```shell
go version
```

# Install Fiber

- install
```shell
go get github.com/gofiber/fiber/v2
```

## Version
- Go: go1.22.5
- Fiber: v2.10.0

## Run
- `main` -> `server.go`
```shell
go run server.go

# or

./run.sh
```

## Golang
- Package Export 규칙
  - 대문자로 시작하면 외부에서 사용 가능
  - 소문자로 시작하면 외부에서 사용 불가능

## Build
- `-o [path]`
```shell
go build -o bin/server
```

## Go 환경변수
```shell
go env
```


## Reference
- https://dev.to/devsmranjan/golang-build-your-first-rest-api-with-fiber-24eh