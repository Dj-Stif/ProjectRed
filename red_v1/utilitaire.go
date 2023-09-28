package red

import "fmt"

// Fonction add inventaire

func (p *Personnage) Addinventory(ItemToAdd string) {

	if len(p.inventaire) >= 10 {
		fmt.Println("Vous ne pouvez pas avoir plus de 10 items dans votre inventaire !")
	} else {
		p.inventaire = append(p.inventaire, ItemToAdd)
		fmt.Printf("Ajout de %s à l'inventaire \n", ItemToAdd)
	}
}

// Fonction remove inventaire
func (p *Personnage) RemoveInventory(itemToRemove string) {
	for i, item := range p.inventaire {
		if item == itemToRemove {
			p.inventaire = append(p.inventaire[:i], p.inventaire[i+1:]...)
			fmt.Printf("%s retirer de l'inventaire \n", itemToRemove)
			return
		}

	}
}

// Fonction Dead
func (p *Personnage) Dead() {
	if p.PV_actuel == 0 {
		fmt.Println("Vous êtes mort !")
		p.PV_actuel = p.PV_max / 2
		p.Affichermenu()
	}
}

func (M *Monster) Monstermort(p *Personnage) {
	if M.monsterPV_actu <= 0 {
		fmt.Println("Vous avez gagné !")
		fmt.Println("Vous avez loot 15 or")
		p.argent += 15
		fmt.Printf("Vous avez %d or\n", p.argent)
		M1.monsterPV_actu = M1.monsterPV_max
		p.Affichermenu()
	}
}

//Fonction Pas assez d'argent
func (p *Personnage) Nomoney(prix int) bool {

	if p.argent >= prix {
		Defilerapide("Vous pouvez acheter l'objet\n")
		p.argent -= prix
		return true
	} else {
		Defilerapide("Vous n'avez pas assez d'argent pour cet item !\n")
		return false
	}
}

//Fonction index scan d'inventaire
func (p *Personnage) Scan(target string) bool {
	for _, scan := range p.inventaire {
		if scan == target {
			return true
		}
	}
	return false
}

//Fonction index scan de skillbook
func (p *Personnage) Scanskill(target string) bool {
	for _, scan := range p.skillbook {
		if scan == target {
			return true
		}
	}
	return false
}
