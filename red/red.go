package piscine

import (
	"fmt"
)

func (p Personnage) DisplayInfo() {
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

func (p Personnage) AccessInventory() {
	var back string
	fmt.Printf("Inventaire : ")
	for cle, val := range p.inventaire {
		fmt.Printf("%d %s\n", val, cle)
	}
	fmt.Println("Appuies sur la touche 'm' pour retourner au menu")
	fmt.Scan(&back)
	if back == "m" {
		p.Menu()
	} else {
		fmt.Println("Mauvaise touche")
		p.AccessInventory()
	}
}

func (p Personnage) Menu() {
	var answer int
	fmt.Println("pour acceder à ton inventaire, tape 'm'. Pour acceder aux informartions de ton personnage, tape 'i'.")
	fmt.Scan(&answer)
	switch answer {
	case 1:
		p.AccessInventory()
	case 2:
		p.DisplayInfo()
	case 3:
		p.Boutique()
	default:
		fmt.Println("Je n'ai pas compris ta requête, peux tu repeter ? ")
	}
}

func (p *Personnage) TakePot(pv int, pvmax int) {
	if p.Pv == p.PvMax {
		fmt.Println("tu as déja toute ta vie!")
	}
	if p.Pv < p.PvMax {
		pv += 50
	}
	if pv > pvmax {
		fmt.Println(pvmax)
	}
	fmt.Println("potion bien utilisée")
}
