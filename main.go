package main

import "fmt"

// func main() {
// 	circ := Circle{
// 		Radius: 7,
// 	}

// 	rect := Rectangle{
// 		weight: 12,
// 		height: 12,
// 	}

// 	fmt.Println(CalculateCircle(circ))
// 	fmt.Println(rect.Area())

// 	helloSay := Sapa{
// 		Kalimat : "rayhan",
// 	}

// 	fmt.Println("hellow", helloSay.SayHello())

// }

// type Shape interface {
// 	Area() float64
// }

// type Sapaan interface{
// 	SayHello()string
// }

// type Circle struct {
// 	Radius float64
// }

// type Rectangle struct{
// 	weight, height float64
// }

// type Sapa struct{
// 	Kalimat string
// }
// //
// func (c Circle) Area() float64 {
// 	return 2 * math.Pi * c.Radius
// }

// func (r Rectangle) Area() float64{
// 	return r.height * r.weight
// }

// func (s Sapa) SayHello()string{
// 	return s.Kalimat
// }

// //

// func CalculateCircle(s Shape) float64{
// 	return s.Area()
// }

// rectangle

type Shape interface{
	Area() int
}

type Rectangle struct{
	Width,Height int
}

// 
func (r Rectangle) Area() int {
	return r.Height * r.Width
}

// func calculate
func CalculateRectangle(s Shape) int {
	return s.Area()
}

// belajar pointer ye nak //

type username struct{
	Name string
}

func UbahUsername(u *username){
	u.Name = "marcello"
}

/////////methode dengan pointer reciver////////////
type User struct{
	Email string
	Username string
	Age int
}

func (u *User) updateEmail(email string){
	u.Email = email
}

func (u *User) updateUsername(username string){
	u.Username = username
}

func main(){
	rectValue := Rectangle{
		Width: 2,
		Height: 2,
	}
	fmt.Println(CalculateRectangle(rectValue))

	user := User{
		Username: "rayahn",
		Email: "kontolodon@gmail.com",
		Age: 20,
	}
	
	user.updateEmail("keren@gmail.com")
	fmt.Println(user)
	user.updateUsername("udah berubah ni")
	fmt.Println(user)
}