Go's app informations
==============

install
===
`go get github.com/axgle/app`

app.Path: Returns the full path of the running executable
as reported by the system. Includes the executable image name.

	package main
	import "github.com/axgle/app"

	func main(){
	  execPath,err:=app.Path()
	  if(err!=nil){
		println("fail get the path")
	   }
	  println(execPath) 
	}
