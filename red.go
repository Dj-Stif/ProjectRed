package main

import "fmt"

type Personnage struct {
	race       string
	classe     string
	niveau     int
	PvMax      int
	Pv         int
	inventaire map[string]int
}

func main() {
	var p1 Personnage
	var p2 Personnage
	var p3 Personnage

	p1.race = "Humain"
	p1.classe = "Guerrier"
	p1.niveau = 1
	p1.PvMax = 100
	p1.Pv = 20
	p1.inventaire = map[string]int{"potion de vie": 1}

	p2.race = "Elfe"
	p2.classe = "Magicienne"
	p2.niveau = 1
	p2.PvMax = 120
	p2.Pv = 25
	p2.inventaire = map[string]int{"potion de vie II": 1}

	p3.race = "Nain"
	p3.classe = "Nain"
	p3.niveau = 1
	p3.PvMax = 90
	p3.Pv = 15
	p3.inventaire = map[string]int{"potion de force": 1}

	p1.displayInfo()
	p2.displayInfo()
	p3.displayInfo()
}

func (p Personnage) displayInfo() {
	fmt.Println("-----------------------")
	fmt.Println("Race :", p.race)
	fmt.Println("Classe :", p.classe)
	fmt.Println("Lvl :", p.niveau)
	fmt.Println("Pv Max :", p.PvMax)
	fmt.Println("Pv :", p.Pv)
	fmt.Printf("Inventaire : ")
	for cle, val := range p.inventaire {
		fmt.Printf("%d %s\n", val, cle)
	}
}
