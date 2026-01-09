package main

import "fmt"

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


// kapan perlu pake pointer
//  1. update state
// pointer use 8 byte
// 2. ketika whenn we want opitmize the memory

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

func Email(e User)string{
	return e.Email
}

func GetAllUser() User{
	return User{
		Username: "hahayy",
	}
}

// pointer dari func lain

func Square (p *int)int{
	*p *= *p
	return *p
}

// pointer jenis lain

type Person struct{
	Name string
	Age int
}

func initPerson() *Person{
	m := Person{
		Name: "nia",
		Age: 17,
	}
	
	return &m
}

type Cafe struct {
	Name string
	Address string
}

func (c *Cafe) UbahNamaCafe(ubah *string)string{
	c.Name = *ubah
	return c.Name
} 

type orang struct{
	namaLengkap string
}

func NewOrang()*orang{
	return &orang{
		namaLengkap: "bodat",
	}
}

func (u *orang) handleChangeName(cn *string)string{
	u.namaLengkap = *cn
	return *cn 
}

func main(){
	x := NewOrang()
	changeNamed := "rayhan keren bangett"
	x.handleChangeName(&changeNamed)
	fmt.Println("sesudah: ",*x)

	cafe := Cafe{
		Name : "rumah rehan",
	}
	fmt.Println("sebelum diubah: ",cafe.Name)

	u := "28"
	res := cafe.UbahNamaCafe(&u)
	fmt.Println("sesudah nama cafe diubah: ",res)

	a := 3
	Square(&a)
	fmt.Println(a)

	fmt.Println(*initPerson())

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
	fmt.Println(GetAllUser())	
}