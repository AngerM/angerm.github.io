package main

import "syscall/js"

func main() {
	document := js.Global().Get("document")
	p := document.Call("createElement", "p")
  p.Set("innerHTML", "Welcome to Matt Anger's Homepage!")
  document.Get("body").Call("appendChild", p)
}
