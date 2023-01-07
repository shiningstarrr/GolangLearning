package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("start")
	const os = runtime.GOOS
	const ar = runtime.GOARCH
	fmt.Println("os:", os)
	fmt.Println("arch:", ar)

}
