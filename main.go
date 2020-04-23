package main

import (
	"member-management/memberControl"
	"member-management/memberView"
)

func main() {
	members := memberControl.NewMembers()

	loop := true
	for loop {
		key := memberView.ShowOpening()
		switch key {
		case 1:
			memberView.ShowMembers(members)
		case 2:
			memberView.AddMember(members)
		case 3:
			memberView.RemoveMember(members)
		case 4:
			loop = false
		}
	}
}
