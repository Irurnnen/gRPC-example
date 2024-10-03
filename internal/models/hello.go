package models

type Hello struct{}

func (h Hello) SayHello(name string) (string, error) {
	return "Hello " + name, nil
}
