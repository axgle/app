package main
import "./.."

func main(){
  execPath,err:=app.Path()
  if(err!=nil){
	println("fail get the path")
   }
  println(execPath) 
}