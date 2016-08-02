package main
import "github.com/axgle/app"

func main(){
  execPath,err:=app.Path()
  if(err!=nil){
	println("fail get the path")
   }
  println(execPath) 
}
