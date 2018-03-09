package employee



type Employee struct {
	Id int
	Name, Position string
}

func EmployeePointerById(id int) *Employee{
	data := map[int]Employee{
		1: {1, "kz", "sz"},
	}
	if e, ok := data[id]; ok {
		return &e
	}
	return nil
}

func EmployeeById(id int) Employee{
	data := map[int]Employee{
		1: {1, "kz", "sz"},
	}
	e, _ := data[id]
	return e
}




/**
eg:

 */