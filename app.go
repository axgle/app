package app

// Returns the full path of the running executable
// as reported by the system. Includes the executable
// image name.
func Path() string{
  path,err:=getExePath();
  if(err!=nil){
    panic("get app.Path fail")
  }
  return path
}
