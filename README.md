# FireDNS
___

###### A just for fun localhost DNS server written in golang.

## Installation
___
> `go get github.com/wisdommatt/firedns`

## Requirements
> * Go 1.9

### Usage
___

1. Import the package `firedns "github.com/wisdommatt/firedns"`
2. Run the program `go run dns.go`
3. On the terminal run `dig meghee.com @localhost -p8090`

## Notes
> The application runs only on port 8090. \
> The only supported domains are meghee.com and timiun.com \
> Checking for an unsupported domain will return an IP of 129.0.1.9 

## Author
[Wisdom Matthew](https://github.com/wisdommatt/)
