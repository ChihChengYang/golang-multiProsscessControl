package main
 

import (
    "fmt" 
)
 
type nvrProcess struct {
}

func (c *nvrProcess ) xProcessInit(){

}

func (c *nvrProcess ) xProcessStart(args []string) error { 
    
    count := len(args)
    for i := 0; i < count; i++{
        fmt.Println( "A->", args[i] )
    }

    return nil 
}

func (c *nvrProcess ) xProcessStop(){
 
}