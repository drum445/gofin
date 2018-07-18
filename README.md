# finance

Methods for performing financial calculations ported from my Crystal version

## Installation

Add this to your source code

```go
import "github.com/drum445/gofin"
```

Run in a terminal in your project root

```bash
$ go get
```

## Examples (shows rounded results)

##### Payment - Returns Float64
```go
# Excel function: =PMT((9.3/100)/12,36,-12995,6000,0)
gofin.PMT(9.3, 36, -12995, 6000, 0) # => 269.92

# Excel function: =PMT((9.3/100)/12,36,-12995,6000,1)
gofin.PMT(9.3, 36, -12995, 6000, 1) # => 267.84
```
##### Present Value - Returns Float64
```go
# Excel function: =PV((9.3/100)/12,36,350.50,1000,0)
gofin.PV(9.3, 36, 350.50, 1000, 0) # => -11731.21

# Excel function: =PV(0,36,350.50,1000,0)
gofin.PV(0, 36, 350.50, 1000, 0) # => -13618.00

# Excel function: =PV((9.3/100)/12,36,350.50,1000,1)
gofin.PV(9.3, 36, 350.50, 1000, 1) # => -11816.26
```
##### NPER - Returns Int32
```go
# Excel function: =NPER(10/100/12,2625.73,-81374.62,0,0)
gofin.NPER(10, 2625.73, -81374.62, 0, 0) # => 36

# Excel function: =NPER(0/100/12,2625.73,-81374.62,0,1)
gofin.NPER(0, 2625.73, -81374.62, 0, 1) # => 31
```
##### Future Value - Returns Float64
```go
# Excel function: =FV((5.8/100)/12,48,2265.63,-133781.21,0)
gofin.FV(5.8, 48, 2265.63, -133781.21, 0) # => 46550.0

# Excel function: =FV((5.8/100)/12,48,2265.63,-133781.21,1)
gofin.FV(5.8, 48, 2265.63, -133781.21, 1) # => 45960.0

# Excel function: =FV((10/100)/12,36,2625.73,-81374.62,0)
gofin.FV(10, 36, 2625.73, -81374.62, 0) # => 0.0
```
##### Rate - Returns Float64
```go
# Excel function: =RATE(60,1600,-75304.59,0) * 12 * 100
gofin.Rate(60, 1600, -75304.59, 0, 0) # => 10.0

# Excel function: =RATE(30,4337.33,-112072.94,0) * 12 * 100
gofin.Rate(30, 4337.33, -112072.94, 0, 0) # => 11.9

# Excel function: =RATE(18,1909.53,-103616.44,77477) * 12 * 100
gofin.Rate(18,1909.53,-103616.44,77477, 0) # => 6.0
```

## Contributing

1. Fork it ( https://github.com/drum445/gofin/fork )
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create a new Pull Request

## Contributors

- [drum445](https://github.com/drum445) ed - creator, maintainer
