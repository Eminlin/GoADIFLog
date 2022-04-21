# GoADIFLog
Parse ADI file like [adif.org](https://adif.org)

Support cabrillo log file like https://www.cqwpx.com/cabrillo.htm

# Usage

```sh
go get -u github.com/Eminlin/GoADIFLog
```
   
```go 
import goadiflog "github.com/Eminlin/GoADIFLog"

func main(){
    logContent, err := goadiflog.Parse(path)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Printf("%+v\n", logContent)
}
```