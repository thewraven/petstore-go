package petstore

type Pet struct {
	ID     string
	Name   string
	Status Status
}

func (p *Pet) GetID() string {
	return p.ID
}

func (p *Pet) SetID(v string) {
	p.ID = v
}

type Order struct {
}

type User struct {
}

type Status int

const (
	Available Status = iota
	Pending
	Sold
)
