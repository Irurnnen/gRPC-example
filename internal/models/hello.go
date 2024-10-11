package models

type Hello struct{}

func (h Hello) SayHello(name string) (string, error) {
	if name == "I want error" {
		return "", ErrExample
	}
	return "Hello " + name, nil
}
