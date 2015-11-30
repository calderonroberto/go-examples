package main

import (
	"fmt"
)

type User struct {
	Name string
}

func (u *User) sayName(speaker string) bool {
	fmt.Printf("Hello %s, my name is %s\n", speaker, u.Name)
	return true
}

type Role interface {
	unlockGate() bool
}

type Admin struct {
	*User
	roleType       string
	SecretPassword string
  entries int32
}

// A method
// Unlock Gate method adding the UnlockGate receiver
//func (s *MyStruct) pointerMethod() { } // method on pointer. Will affect the receiver
//func (s MyStruct)  valueMethod()   { } // method on value. Will not affect the receiver, and respects the immutable nature of the object (it gets a copy)
// this method has the same signature [UnlockGate() bool] specified by the Role interface, golang infers this is then a Role
func (a *Admin) unlockGate() bool {
  a.entries++
	fmt.Printf("unlocking gate as %s\n", a.Name)
	return true
}

//if a variable has an interface type then we call methods named in the interface. Here this generic function takes advantage of this to work on any role
func unlock(r Role) {
	r.unlockGate()
}

func main() {

	//using pointer makes it cheap (method will not create a new receiver, but will use the pointer); and allowing it to modify the original object.
	b := &Admin{&User{"Pancho"}, "administrator", "pa$$word", 0}
	b.sayName("Roberto")

	b.unlockGate()
	unlock(b)

  fmt.Printf("Total entries %d\n", b.entries)

}
