package red

import "fmt"

// Fonction add inventaire

func (p *personnage) addinventory(ItemToAdd string) {
	if len(p.inventaire) >= 10 {
		fmt.Println("Vous ne pouvez pas avoir plus de 10 items dans votre inventaire !")
	} else {
		p.inventaire = append(p.inventaire, ItemToAdd)
		fmt.Printf("Ajout de %s à l'inventaire \n", ItemToAdd)
	}
}

// Fonction remove inventaire
func (p *personnage) removeInventory(itemToRemove string) {
	for i, item := range p.inventaire {
		if item == itemToRemove {
			p.inventaire = append(p.inventaire[:i], p.inventaire[i+1:]...)
			return
		}

	}
}

// Fonction Dead
func (p *personnage) Dead() {
	if p.PV_actuel == 0 {
		fmt.Println("Vous êtes mort !")
		p.PV_actuel = p.PV_max / 2
		fmt.Println("Vous avez été ressuscité avec 50% hp")
	}
}

// Fonction est ce que c'est dans l'inventaire
func (p *personnage) IsInSkill(skl string) bool {
	for _, skills := range p.skills {
		if skl == skills {
			return true
		}
	}
	return false
}

//Fonction est ce que c'est dans l'inventaire
func (p *personnage) IsInInventory(val string) bool {
	for char := range p.inventaire {
		if val == char {
			return true
		}
	}
	return false
}
