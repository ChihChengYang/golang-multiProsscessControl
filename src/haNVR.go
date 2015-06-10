package main
  
import (
    "fmt"
    "net/http"
    "os"  
)
 
func main() {
 
    var cp childProcess 
    cp.childProcessInit()
    cp.params = append(cp.params, "10.62.8.102/mpeg4")
    cp.params = append(cp.params, "admin")
    cp.params = append(cp.params, "123456")
 
    port , _ := cp.childProcessStart()
   
    fmt.Println("port .....", port)

    var num float32
    fmt.Scanln(&num)
    cp.childProcessStop(port)

    fmt.Scanln(&num)
    cp.childxProcessCommutation(port)
 

    http.ListenAndServe(":11002", nil)
}
 
