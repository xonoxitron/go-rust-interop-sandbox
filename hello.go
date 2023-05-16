// main.go
package main

// #cgo LDFLAGS: -L./ -lhello
// void greet();
import "C"

func main() {
	C.greet()
}
