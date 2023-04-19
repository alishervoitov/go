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

// package main

// import "fmt"

// func main() {
// 	var name string
// 	fmt.Println("What's your name")
// 	fmt.Scan(&name)
// 	fmt.Println("Hello " + " " + name)
// }

// package main

// import "fmt"

// func main() {
// 	num := -3
// 	if num > 0 {
// 		fmt.Println("Number is greater than 0")
// 	} else if num < 0 {
// 		fmt.Println("Number is little than 0")
// 	} else if num == 3 {
// 		fmt.Println("Number equals 3 !!!")
// 	}
// }

package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	var c float64

	fmt.Println("Enter the number a")
	fmt.Scan(&a)

	fmt.Println("Enter the number b")
	fmt.Scan(&b)

	fmt.Println("Enter the number c")
	fmt.Scan(&c)

	D := b*b - 4*a*c

	if D > 0 {
		var x1 float64
		var x2 float64

		x1 = (-b + math.Sqrt(D)) / 2 * a
		x2 = (-b - math.Sqrt(D)) / 2 * a
		fmt.Println(fmt.Sprint(x1))
		fmt.Println(fmt.Sprint(x2))
	} else if D == 0 {
		var x float64

		x = (-b + math.Sqrt(D)) / 2 * a

		fmt.Println(x)
	} else if D < 0 {
		fmt.Println("You can't solve this issue")
	}
}
