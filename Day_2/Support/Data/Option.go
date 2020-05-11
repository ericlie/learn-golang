package Data

type Action func(input string)

type Option struct {
	Name string
	Description string
	Arguments []string
	Action Action
}