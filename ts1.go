package lab1


func  Task1(inpt string) (string, error) {
 op1:=""
 op1fl := false 
 op2:=""
 for i :=0; i<len(inpt);i++{

  if string(inpt[i])=="*" || string(inpt[i])=="+"{
    op1 =string(inpt[i])+op2+op1
    op2=""
   } else if string(inpt[i])=="-" ||string(inpt[i])=="/" {
    op1 =string(inpt[i])+op1+op2
    op2=""
    }else if(string(inpt[i])!=" "&&op1fl==false){
    	if(op1==""){
    		op1=" "
    	}
     op1+=string(inpt[i])
   }else if(string(inpt[i])!=" "&&op1fl==true){
    op2+=string(inpt[i])
   }else if(string(inpt[i])==" "){
    if(op1fl==false){
     op1+=string(inpt[i])
     op1fl = true
     }else{
      op2+=string(inpt[i])
     }
   }
 }

 return op1[:len(op1)-1],nil
}
