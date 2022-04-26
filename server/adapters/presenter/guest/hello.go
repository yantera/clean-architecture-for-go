package guest

type Hello struct {
	Message1 string `json:"message1"`
	Message2 string `json:"message2"`
}

func HelloWorld() Hello {
	hello := Hello{}

	hello.Message1 = "Hello"
	hello.Message2 = "World"
	return hello
}
