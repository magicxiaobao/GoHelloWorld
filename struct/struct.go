package main

import "fmt"

type Programmer struct {
	name      string
	languages []string
	hobbies   []string
}

func main() {

	var xiaobao = Programmer{"xiaobao", []string{"java", "go"}, []string{"sing", "video games"}}
	fmt.Println(xiaobao)
	fmt.Println("xiaobao's hobbies", xiaobao.hobbies)
	var xiaobaoCopy = xiaobao
	xiaobao.languages = append(xiaobao.languages, "vue")
	fmt.Println(xiaobaoCopy, xiaobao)
	var suisui Programmer
	suisui.name = "suisui"
	fmt.Println("suisui", suisui)
	var guoguo = Programmer{
		languages: []string{"python"},
		name:      "guoguo",
	}
	fmt.Println("guoguo", guoguo)
	programmerPointerByNew := new(Programmer)
	fmt.Println(programmerPointerByNew)
	programmerPointerByNew.languages = []string{"java", "vue", "go", "python"}
	programmerPointerByNew.name = "chenlaoxie"
	programmerPointerByNew.hobbies = []string{"sing", "video game"}
	fmt.Println(*programmerPointerByNew)
	// structs are value types
}
