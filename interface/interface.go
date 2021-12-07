package main

import "fmt"

type play interface {
	videoGame() string
	computerGame() string
}
type godOfWar struct {
	name string
}

func (g *godOfWar) videoGame() string {
	return g.name + " can play in video game"
}

func (g *godOfWar) computerGame() string {
	return g.name + " can play in computer game"
}

func someOnePlay(p play) {
	fmt.Println(p.videoGame() + " and " + p.computerGame())
}

func main() {

	godOfWar := godOfWar{name: "kuiduosi"}
	fmt.Println(godOfWar.computerGame())
	fmt.Println(godOfWar.videoGame())
	someOnePlay(&godOfWar)
}
