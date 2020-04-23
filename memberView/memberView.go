package memberView

import (
	"fmt"
	"member-management/memberControl"
)

type MemberView struct {}

func NewMemberView() *MemberView {
	return &MemberView{}
}

type MemberShowInterface interface {
	ShowOpening () int
	ShowMembers(mc *memberControl.MemberControl)
}

type MemberAddInterface interface {
	AddMember(mc *memberControl.MemberControl)
}


type MemberRemoveInterface interface {
	RemoveMember(mc *memberControl.MemberControl)
}

func (mv *MemberView)ShowOpening () int {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("-------------------------Member Management-------------------------")
	fmt.Println("                         1. Show all members")
	fmt.Println("                         2. Add a member")
	fmt.Println("                         3. Remove a member")
	fmt.Println("                         4. Exit")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Print("Select a number: ")
	key := 0
	fmt.Scanln(&key)
	return key
}

func (mv *MemberView)AddMember(mc *memberControl.MemberControl) {
	fmt.Println()
	fmt.Print("name: ")
	var name string
	fmt.Scanln(&name)

	fmt.Println()
	fmt.Print("phone: ")
	var phone string
	fmt.Scanln(&phone)

	fmt.Println()
	fmt.Print("age: ")
	var age int
	fmt.Scanln(&age)

	fmt.Println()
	fmt.Print("gender: ")
	var gender string
	fmt.Scanln(&gender)

	mc.AddMember(name, phone, age, gender)
	fmt.Println("Success to add a member")
	fmt.Println()
}

func (mv *MemberView)ShowMembers(mc *memberControl.MemberControl) {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("\t id \t name \t phone \t age \t gender ")
	for _, item := range mc.Members {
		println()
		fmt.Printf("\t %v \t %v \t %v \t %v \t %v ", item.Id, item.Name, item.Phone, item.Age, item.Gender)
	}
	println()
	println("Input q to back to memu")
	var key string
	for key != "q" {
		fmt.Scanln(&key)
	}
}

func (mv *MemberView)RemoveMember(mc *memberControl.MemberControl) {
	fmt.Println()
	fmt.Print("id: ")
	var id int
	fmt.Scanln(&id)


	mc.RemoveMember(id)
	fmt.Println("Success to remove the member")
	fmt.Println()
}