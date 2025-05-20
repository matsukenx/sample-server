package usecase

type HelloUsecase interface {
	SayHello(name string) string
}

type helloUsecase struct{}

func NewHelloUsecase() HelloUsecase {
	return &helloUsecase{}
}

func (u *helloUsecase) SayHello(name string) string {
	if name == "" {
		name = "World"
	}

	return "Hello, " + name + "!"
}
