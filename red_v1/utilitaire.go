package red

import "fmt"

// Fonction add inventaire
func (p *Personnage) AddInventory() {
	p.inventaire = append(p.inventaire)
	fmt.Println("Ajout à l'inventaire")

}

// Fonction remove inventaire
func (p *Personnage) RemoveInventory(itemname string) {
	for i, itemname := range p.inventaire {
		p.inventaire = append(p.inventaire[:itemname[i]], p.inventaire[itemname[i]:]...)
	}
}

// Fonction Dead
func (p *Personnage) Dead() {
	if p.PV_actuel == 0 {
		fmt.Println("Vous êtes mort !")
		p.PV_actuel = p.PV_max / 2
		fmt.Println("Vous avez été ressuscité avec 50% hp")
	}
}

// Fonction est ce que c'est dans l'inventaire
func (p *Personnage) IsInSkill(skl string) bool {
	for _, skills := range p.skills {
		if skl == skills {
			return true
		}
	}
	return false
}

//Fonction est ce que c'est dans l'inventaire
func (p *Personnage) IsInInventory(val string) bool {
	for char := range p.inventaire {
		if val == char {
			return true
		}
	}
	return false
}
