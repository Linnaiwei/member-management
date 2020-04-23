package main

import (
	"member-management/memberControl"
	"member-management/memberView"
)


func main() {
	mc := memberControl.NewMemberControl()
	mv := memberView.NewMemberView()
	loop := true
	for loop {
		key := mv.ShowOpening()
		switch key {
		case 1:
			mv.ShowMembers(mc)
		case 2:
			mv.AddMember(mc)
		case 3:
			mv.RemoveMember(mc)
		case 4:
			loop = false
		}
	}
}
