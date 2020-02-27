package main

import(
       "fmt"
        lab1 "github.com/project"
)

func main() {
    res,_ := lab1.Task1("4 2-3*5+")
    fmt.Println(res)
    
    
    fmt.Println(buildVersion)
}
/*
func Task1(inpt string){
	op1:="a"
	var op2 string
	for i :=0; i<len(inpt);i++{
		if string(inpt[i])=="*" ||
		 string(inpt[i])=="+"{
		 	op1 =string(inpt[i])+op2+op1
		 }else if string(inpt[i])=="-" ||
		 string(inpt[i])=="/" {
		 	op1 =string(inpt[i])+op1+op2
		 	}else if string(inpt[i])==" " {
		 		op1 +=" "
		 	}else if op1=="a"{
		 		op1=string(inpt[i])
		 	}else{
		 		op2=string(inpt[i])
		 	}
	}
	fmt.Println(op1)
}*/