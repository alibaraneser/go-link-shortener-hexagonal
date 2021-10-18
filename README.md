# 🚀  go-link-shortener-hexagonal

#### Link Shortener which is using Hexagonal Architecture that compliant with clean code guidelines in Go.

```bash
├── README.md
├── cmd
│   └── http.go
├── go.mod
├── go.sum
├── internal
│   ├── adapters
│   │   ├── repository
│   │   │   └── redis
│   │   │       └── repository.go
│   │   └── serializer
│   │       └── json
│   │           └── serializer.go
│   ├── core
│   │   ├── business
│   │   │   └── logic.go
│   │   ├── domain
│   │   │   ├── errors.go
│   │   │   └── redirect.go
│   │   └── ports
│   │       ├── repository.go
│   │       ├── serializer.go
│   │       └── service.go
│   └── factory
│       └── repository.go
├── main.go
├── pkg
│   ├── api
│   │   └── server.go
│   └── generator
│       └── random.go
```


