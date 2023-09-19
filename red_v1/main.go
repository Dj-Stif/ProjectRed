package red

import (
	"fmt"
	"time"
)

// Fonction TakePot
func (p *personnage) takepot() {

	diffPV := p.PV_actuel + 50 - p.PV_max
	if p.PV_actuel == p.PV_max {
		fmt.Println("Vous avez déjà le maximum de PV")
	}

	if diffPV < 0 {
		p.PV_actuel += 50
		fmt.Println("Vous avez utilisé une potion, vous avez :", p.PV_actuel, "PV")
		p.removeinventory("Potion")

	}
	var choixpotion int
	if diffPV > 0 {
		fmt.Println("Votre potion ne sera pas utilisé complètement, voulez-vous quand même l'utiliser ?")
		fmt.Println("1 - Oui")
		fmt.Println("2 - Non")
		fmt.Scan(&choixpotion)

		switch choixpotion {
		case 1:
			p.PV_actuel = p.PV_max
			fmt.Println("Vous avez utilisé une potion, vous avez :", p.PV_actuel, "PV")
			p.removeinventory("Potion")

		case 2:
			p.accessinventory()
		}
	}
}

// Fonction potion de poison
func (p *personnage) poisonpot() {
	var durée int = 3
	var degatspoison int = 10

	fmt.Println("Vous avez été empoisonné")
	for i := 0; i < durée; i++ {
		p.PV_actuel -= degatspoison
		fmt.Println("Vos PV : ", p.PV_actuel)
		time.Sleep(1 * 3)
	}
	fmt.Println("Le poison a cessé de faire effet")
}
