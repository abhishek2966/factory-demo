package resource

const (
	BackendTeam  = "backend"
	FrontendTeam = "frontend"
	InfraTeam    = "infra"
)

type Employee struct {
	Name  string
	Email string
	Team  string
}

func NewEmployeeGenerator(team string) func(string, string) Employee {
	f := func(name, email string) Employee {
		return Employee{
			Name:  name,
			Email: email,
			Team:  team,
		}
	}

	return f
}
