#!/bin/sh
 
go build xProcess.go nvr.go 
go build haNVR.go childProcessesHandle.go 
mv ./xProcess ../bin/
mv ./haNVR ../bin/