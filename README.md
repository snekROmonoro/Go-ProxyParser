# Go-ProxyParser
Library in Golang to parse proxy URL details

## Installation

```sh
go get github.com/snekROmonoro/Go-ProxyParser
```

## Usage

```go
import (
  proxyparser "github.com/snekROmonoro/Go-ProxyParser"
)

...

var proxyStr = "http://username:password@host:port"
var ProxyData *proxyparser.ProxyData
var parsedProxySuccessfully = false

ProxyData, parsedProxySuccessfully = proxyparser.GetProxyData(proxyStr)
```
