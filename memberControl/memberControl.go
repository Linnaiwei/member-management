package memberControl

import "member-management/member"

type Members []member.Member

func (members *Members) AddMember (name string, phone string, age int, gender string) {
	id := 0
	if len(*members) > 0 {
		id = (*members)[len(*members)-1].Id + 1
	}
	*members = append(*members, *member.NewMember(id, name, phone, age, gender))
}

func (members *Members) RemoveMember (id int) {
	for index, item := range *members {
		if item.Id == id {
			*members = append((*members)[:index], (*members)[index+1:]...)
			break
		}
	}
}

func NewMembers () *Members {
	return new(Members)
}
