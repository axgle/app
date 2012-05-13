package app

func getExePath() (exePath string, err error) {
	return os.Readlink(`/proc/self/exe`)
}
