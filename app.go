// Go's app informations
package app

// Returns the full path of the running executable
// as reported by the system. Includes the executable
// image name.
func Path()  (string, error) {
  return getExePath(); 
}
