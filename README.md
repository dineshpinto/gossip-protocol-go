[![codecov](https://codecov.io/gh/dineshpinto/gossip-protocol-go/graph/badge.svg?token=VJZB5A3A91)](https://codecov.io/gh/dineshpinto/gossip-protocol-go)
[![Go](https://github.com/dineshpinto/gossip-protocol-go/actions/workflows/go.yml/badge.svg)](https://github.com/dineshpinto/gossip-protocol-go/actions/workflows/go.yml)

# gossip-protocol-go

A synchronous gossip protocol with Byzantine nodes implemented in Golang.

## Installation

```bash
go get -u github.com/dineshpinto/gossip-protocol-go
```

## Run 
    
```bash
go run main.go
```

## Benchmark

```bash
go test -bench=. -run=^# -benchtime=50x ./... 
```

| Param           | Value |
|-----------------|-------|
| num_non_sample  | 1000  |
| num_honest      | 6     | 
| num_adversarial | 4     |
| num_peers       | 6     |
| cycles          | 200   |

```
goos: darwin
goarch: arm64
pkg: github.com/dineshpinto/gossip-protocol-go/node
BenchmarkCreateNodes-8                        50            160448 ns/op
BenchmarkConnectNodesToRandomPeers-8          50           6141139 ns/op
BenchmarkEvolveState-8                        50        1424933124 ns/op
```