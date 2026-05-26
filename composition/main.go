package main

type Engine struct {
	Type string
}

func (e Engine) Start() {
	println("Starting the", e.Type, "engine")
}

type Car struct {
	Model string
	Engine
}

func (c Car) Drive() {
	println("Driving the", c.Model)
	c.Engine.Start()
}

func main() {
	car := Car{
		Model: "Sedan",
		Engine: Engine{
			Type: "V8",
		},
	}
	car.Drive()
}
