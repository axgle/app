// Go's app informations
package app

// Returns the full path of the running executable
// as reported by the system. Includes the executable
// image name.
func Path()  (string, error) {
  return getExePath(); 
}


func MustPath() (string) {
  path,err:=getExePath(); 
  if err!=nil {
   panic("app.MustPath error")
  }
  return path
}