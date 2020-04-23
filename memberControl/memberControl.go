package memberControl

import "member-management/member"

type MemberControl struct {
	Members []member.Member
}


func (this *MemberControl) AddMember (name string, phone string, age int, gender string) {
	id := 0
	if len(this.Members) > 0 {
		id = (this.Members)[len(this.Members)-1].Id + 1
	}
	this.Members = append(this.Members, *(member.NewMember(id, name, phone, age, gender)))
}

func (this *MemberControl) RemoveMember (id int) {
	for index, item := range this.Members {
		if item.Id == id {
			this.Members = append((this.Members)[:index], (this.Members)[index+1:]...)
			break
		}
	}
}

func NewMemberControl () *MemberControl {
	return new(MemberControl)
}
