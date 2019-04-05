package exam

type exam struct {
	id int
	name string
}

func (e *exam) GetName() string {

	return "浙江工商大学"

}

func (e *exam) GetId() int {
	return 23
}

