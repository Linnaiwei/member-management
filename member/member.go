package member

type Member struct {
	Id     int
	Name   string
	Phone  string
	Age    int
	Gender string
}

func NewMember(id int, name string, phone string, age int, gender string) *Member {
	return &Member{Id: id, Name: name, Phone: phone, Age: age, Gender: gender}
}

