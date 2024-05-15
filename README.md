# Serial Port
serialport is a small library which allows you to send or received data from hardware
```shell
go get github.com/go-pkg-utils/serialport
```
For example:
------------
```go
func main() {
    sp, err := serialport.Open("COM1", 9600, 8, 1)
    if err != nil {
        fmt.Println(err)
    } else {
        sp.Received(0x0a, func(data []byte) {
            fmt.Printf("%s", data)
        })
    }

    fmt.Scanln()
}
```
