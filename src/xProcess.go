package main

import (
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "strconv"
)

type xProcessInterface interface {
    xProcessInit()
    xProcessStart(Args []string) error 
    xProcessStop() 
}

type xProcess struct {
	xpi xProcessInterface 
}


func (x *xProcess) xOpen( args []string ) error{
 
    count := len(args)
    for i := 0; i < count; i++{
        fmt.Println( args[i] )
    }
    
    return nil
}

func (x *xProcess) xClose() {
    fmt.Println("xClose!!!!!!!!")
 
//---------------------    
 //    timer := time.NewTimer(time.Second * 2)
  //   <- timer.C
   //  println("Timer expired")
    os.Exit(0)
//---------------------    
}

func (x *xProcess) xComm(w http.ResponseWriter, r *http.Request) { 
  
    r.ParseForm()
    if r.Method == "GET" { 
        fmt.Println("nvrComm GET", os.Getppid() )
    }else if r.Method == "POST" { 
        
        if id := r.FormValue("id"); id != strconv.Itoa(os.Getppid()) {
            fmt.Fprintf(w, "ERR")
        }else{
            fmt.Println("OK ", id )            
            if clo := r.FormValue("close"); clo == "yes"{
    
                go x.xClose() 
            }
             
        }
 
    }
 
    fmt.Fprintf(w, "OK")  
}
 
func main() {
  
    var x xProcess
    argsport := len(os.Args) - 1
    port,_ := strconv.Atoi(os.Args[argsport])

    if port <= 0 {
        fmt.Println("xProcess Port ERROR!!.....",os.Args[argsport] )
        return
    }

    x.xOpen(os.Args)

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    signal.Notify(c, syscall.SIGTERM)
    go func() {
        <-c
        x.xClose()
        os.Exit(0)
    }()
 
    defer x.xClose() 
 
    fmt.Println(" http.HandleFunc.....")

    http.HandleFunc("/",  x.xComm)
    http.ListenAndServe(":"+os.Args[argsport], nil)

    fmt.Println(" http.HandleFunc.....")
}
 