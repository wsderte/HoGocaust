package lab1


func  Task1(inpt string) ( string, error ) {
	op1:="a"
	var op2 string
	for i :=0; i<len(inpt);i++{
		if string(inpt[i])=="*" ||
		 string(inpt[i])=="+"{
		 	op1 =string(inpt[i])+op2+" "+op1
		 }else if string(inpt[i])=="-" ||
		 string(inpt[i])=="/" {
		 	op1 =string(inpt[i])+op1+" "+op2
		 	}else if string(inpt[i])==" " {
		 	}else if op1=="a"{
		 		op1=string(inpt[i])
		 	}else{
		 		op2=string(inpt[i])
		 	}
	}
	return op1,nil
}