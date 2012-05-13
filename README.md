Go's app informations
==============
app.Path: Returns the full path of the running executable
as reported by the system. Includes the executable image name.

	package main
	import "github.com/axgle/app"

	func main(){  
	  println(app.Path())
	}