// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, world")
// }

// package main

// import "fmt"

// func main() {
// 	var age int = 8
// 	fmt.Println(age)
// }

// package main

// import "fmt"

// func main() {
// 	var number float64 = 222.222
// 	fmt.Println(number)
// }

// package main

// import "fmt"

// func main() {
// 	age := 14.444333
// 	fmt.Println(age)
// }

// package main

// import "fmt"

// func main() {
// 	var name string = "Alisher"
// 	fmt.Println(name)
// 	fmt.Println(len(name))
// }

package main

import "fmt"

func main() {
	var name string
	fmt.Println("What's your name")
	fmt.Scan(&name)
	fmt.Println("Hello" + name)
}
