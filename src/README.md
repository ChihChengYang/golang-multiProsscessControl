
golang-multiProsscessControl

================

A golang program for multi-prosscess control  


parent process port 11002 haNVR.go 

    childProcessesHandle.go <interface> <---port 11003--> xProcess.go arouse Child Process 1  
                                        <---port 11004--> xProcess.go arouse Child Process 2   
                                        <---port 11005--> xProcess.go arouse Child Process 3  
                                        ......
Usage
================

###Install

./build_haNVR.sh

###Examples

1.  Initial one child Process

    var cp childProcess 
    cp.childProcessInit() 

2.  Set params for xProcess (e.g. your own program)  

    cp.params = append(cp.params, "10.62.8.102/mpeg4")
    cp.params = append(cp.params, "admin")
    ...

3.  Start the child Process 
    port , _ := cp.childProcessStart()

4.  Commutation with the child Process ( unfinished implement...)   
    cp.childxProcessCommutation(port)

5.  Stop the child Process 
    cp.childProcessStop(port)    

 
Document
================

 