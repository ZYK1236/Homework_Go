package logMsg

import "fmt"

func LogSuccessMsg(path string, method string) {
	fmt.Println("path:", path, "✅ ------>", method)
}

func LogErrorMsg(path string, method string) {
	fmt.Println("path:", path, "❌ ------>", method)
}
