package app
import "os"
func getExePath() (exePath string, err error) {
	return os.Readlink(`/proc/self/exe`)
}