package main

import pr "piscine/red"

func main() {
	var p1 pr.Personnage
	p1.Init("Humain", "Guerrier")
	p1.Menu()
	p1.DisplayInfo()
}
