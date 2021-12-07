package main

import "fmt"

type Programmer struct {
	name      string
	languages []string
	hobbies   []string
}

func (p Programmer) Learn(subject string) string {

	p.hobbies = append(p.hobbies, "biubiubiu")
	return p.name + " is learning " + subject
}

func learn(p *Programmer, subject string) string {

	p.hobbies = append(p.hobbies, "swimming")
	return p.name + " is learning " + subject
}

func main() {

	xiaobao := Programmer{name: "xiaobao"}
	fmt.Println(xiaobao)
	println((&xiaobao).Learn("go"))
	fmt.Println(xiaobao)
	learn(&xiaobao, "vue3")
	fmt.Println(xiaobao)
}
