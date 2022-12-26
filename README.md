# speedtest
The small GO library that tests the download and upload speeds by using Ookla's https://www.speedtest.net/ and Netflix's https://fast.com/.

## Dependencies
- [https://github.com/showwin/speedtest-go](https://github.com/showwin/speedtest-go) -  pure Go API to Test 
Internet Speed using [speedtest.net](speedtest.net)
- [https://github.com/adhocore/fast](https://github.com/adhocore/fast) - Go command line tool to check internet speed 
on [fast.com](fast.com).

## Usage 
```
go get github.com/shchiv/speedtest
```

### Speedtest API

#### Netflix (fast.com) example
```go
func main() {
    speed, err := speedtest.Start(speedtest.Netflix())
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Netflix Download: %v Upload: %v", speed.Down, speed.Up)
}
```

Output:
```
Netflix    
    Download: 67 Mbps    
      Upload: 83 Mbps
```           

#### Ookla (speedtest.next) example
```go
func main() {
    speed, err := speedtest.Start(speedtest.Ookla())
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Netflix Download: %v Upload: %v", speed.Down, speed.Up)
}
```

Output:
```
Ookla    
    Download: 65 Mbps    
      Upload: 89 Mbps
```          

## Test Coverage
```
go test ./... --cover -covermode=count
ok      github.com/shchiv/speedtest/cmd/speedtest       

0.209s  coverage: 100.0% of statements

```

## Benchmark
```
go test ./... -bench=. -run=^#    
goos: darwin
goarch: arm64
pkg: github.com/shchiv/speedtest/cmd/speedtest
BenchmarkOokla-10              1        23433669000 ns/op
BenchmarkNetflix-10            1        38330611958 ns/op

```