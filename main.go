package main

import "fmt"

// func main() {
// 	r := config.InitRoute()
// 	r.Run(":9999")
// }
func decorator(f func(s string)) func(s string, more string) {
	return func(s string, more string) {
		fmt.Println("Started")
		f(s)
		fmt.Println("Done")
	}
}
func Hello(s string) {
	fmt.Println(s)
}
func main() {
	decorator(Hello)("Hello, World!", "liuyong")
}
