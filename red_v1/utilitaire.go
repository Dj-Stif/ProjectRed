package red

import "fmt"

// Fonction add inventaire
func (p *personnage) AddInventory() {
	p.inventaire = append(p.inventaire)
	fmt.Println("Ajout à l'inventaire")

}

// Fonction remove inventaire
func (p *personnage) RemoveInventory(itemname string) {
	for i, itemname := range p.inventaire {
		p.inventaire = append(p.inventaire[:itemname[i]], p.inventaire[itemname[i]:]...)
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
func (p personnage) IsInSkill() {
	for skl, pos := range p.skills {
		fmt.Println("Tu connais déjà ce skill")
	}
}

//Fonction est ce que c'est dans l'inventaire
func (p personnage) IsInInventory() {
	for sb, val := range p.inventaire {
		fmt.Println("Tu as déjà ce SpellBook")
	}
}
