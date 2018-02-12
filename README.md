# Nadire - the Random Data Server (WIP)
[![Go Report Card](https://goreportcard.com/badge/github.com/Peter42/nadire)](https://goreportcard.com/report/github.com/Peter42/nadire)

**Disclaimer**
This is one of my first times using Go.
Don't expect good (or even average) code quality and feel free to open bug on bad code smell.

### Download
Get the executable for your OS and Architecture from the [Releases](https://github.com/Peter42/nadire/releases) section of this repo.

### Usage
Start the executable, send as many request as you want.

Port:
- the first argument will be used as port
- default port is 4321

Examples:
- [?size=1k](http://localhost:4321/?size=1k) returns 1 KB of random data.
- [?delay=10000&size=1k](http://localhost:4321/?delay=10000&size=1k) returns 1 KB of random data after a delay of (about) 10 seconds.
- [?size=100M](http://localhost:4321/?size=100M) returns 100 MB of random data.
- [?size=512](http://localhost:4321/?size=512) returns 512 Byte of random data.

### Build
1. Download and install Go.
2. Clone repo
3. Navigate to repo with your terminal
4. Run `go build`. This will build Nadire for your operating system and processor architecture.

### Name
As the Server generated random data it's name was chosen by random via [behindthename.com](https://www.behindthename.com/random/random.php?number=1&gender=both&surname=&all=yes)
