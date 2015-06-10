
golang-multiProsscessControl

================

A golang program for multi-prosscess control  


parent process port 11002 haNVR.go 

    childProcessesHandle.go <os.StartProcess> <---port 11003--> arouse xProcess.go <interface> Child Process 1 (nvr.go) 
                                              <---port 11004--> arouse xProcess.go <interface> Child Process 2 (nvr.go)   
                                              <---port 11005--> arouse xProcess.go <interface> Child Process 3 (nvr.go)  
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
    cp.childProcessCommutation(port)

5.  Stop the child Process 
    cp.childProcessStop(port)    

 
Document
================

 