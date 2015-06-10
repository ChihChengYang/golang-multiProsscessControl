package main
 
import (
    "fmt"
    "net/http"
    "os"  
    "time"
    "io/ioutil"
    "strings"
    "strconv"
    "errors"
)

var (
    Err_ChildProcessNonValidPort   = errors.New("child Process Get Non-Valid Port")
)

const gPortNumbers = 100

const gPortFrom = 11005
 
type childxProcess struct {
    xIPexited  chan string
    p *os.Process 
    params []string
    port int
}

type childProcessPort struct {
    port int
    flag bool
}
 
var gChildProcessPort map[int]childProcessPort

var gxProcessMap map[int]childxProcess   

type childProcess struct {
    params []string
} 
 
func (xprocess *childxProcess) xProcessHandle(){
    
    exitflag := false
    for {
        select {
            case  ip := <-xprocess.xIPexited:
                fmt.Println("nvrProcessHandle: ",ip)
                exitflag = true
                break
            default:
        }
 
        if(exitflag){
            break
        }

        time.Sleep(10 * time.Millisecond)
    }
    fmt.Println("xProcessHandle Exit")
}
 

func (xprocess *childxProcess) xProcessStart(port int) {

    if(len(xprocess.params)<=0){
        println("Params Count Fail:" )
        return
    }

    var procAttr os.ProcAttr
    procAttr.Files = []*os.File{nil, os.Stdout, os.Stderr}
 
    var perr error 
 
    xprocess.p, perr = os.StartProcess("xProcess",  
        xprocess.params, 
        &procAttr)
    if perr != nil {
        println("start process failed:" + perr.Error())
        return
    }

    xx := gxProcessMap[port]
    xx.p = xprocess.p
    gxProcessMap[port] = xx

    processState, err := xprocess.p.Wait()
    if err != nil {
        println("Wait Error:" + err.Error())
    }

    println("processState:" , processState.Pid())
    
    xprocess.xIPexited <-  xprocess.params[0]

}

func xProcessGo( port int ) {
 
    xp := gxProcessMap[port]
    go xp.xProcessStart(port)
    go xp.xProcessHandle()
}

func xProcessKill( port int ) {   
    xp := gxProcessMap[port]
    if(xp.p!=nil){
        xp.p.Signal(os.Interrupt )
        xp.xIPexited <- xp.params[0]
    }
    
    fmt.Println("xProcessKill ", xp.p)
}
 
func childProcessGetValidPort() (port int) {
    for _, value := range gChildProcessPort {        
        if value.flag == false{
            return value.port  
        }
    }    
    return -1
}

func childProcessValidPortInit() {
    gChildProcessPort = make(map[int]childProcessPort)
    for i := 0; i < gPortNumbers; i++{
        var c childProcessPort
        c.port = gPortFrom + i;
        c.flag = false
        gChildProcessPort[i] = c
    }
}
 
func (cp *childProcess) childProcessInit(){
    childProcessValidPortInit()
    gxProcessMap = make(map[int]childxProcess) 
}
 
func (cp *childProcess) childProcessStart() (int , error) {  

    port := childProcessGetValidPort()
    if port == -1{
        return port,Err_ChildProcessNonValidPort
    }
 
    var xp childxProcess
 
    xp.params = cp.params
    xp.params = append(xp.params, strconv.Itoa(port))
 
    xp.port =  port
    xp.xIPexited = make(chan string)
    xp.p = nil
 
    gxProcessMap[xp.port] = xp
 
    xProcessGo( xp.port )

    return port,nil
}

func (cp *childProcess) childProcessStop(port int){
    xProcessKill(port)
    delete(gxProcessMap, port)
}

func (cp *childProcess) childProcessCommutation(port int) {
    
    if port <= 0 {
        fmt.Println("childProcessCommutation Port ERROR!!.....")
        return
    }

    url := "http://127.0.0.1:"+ strconv.Itoa(port)

    timeout := time.Duration(5 * time.Second)
    client := http.Client{
        Timeout: timeout,
    }
   
 
    req, err := http.NewRequest("POST", url, strings.NewReader("id="+strconv.Itoa(os.Getpid())+"&close=no"))
    //req.Header.Set("X-Custom-Header", "myvalue")
    //req.Header.Set("Content-Type", "text/plain")
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
    response, err := client.Do(req)
 
    if err != nil {
        fmt.Printf("====>%s", err)
        //os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("-----> %s", err)
          //  os.Exit(1)
        }
        fmt.Printf(">>>>%s\n", string(contents))
    }
 
}