package red

import "fmt"

//FORGERON

func (p *Personnage) Accesforgeron() {
	Defilerapide("Bienvenue chez le forgeron : \n")
	Defilerapide("1. Chapeau de l'aventurier (1 Plume de Corbeau et 1 Cuir de Sanglier)\n")
	Defilerapide("2. Tunique de l'aventurier (2 Fourrures de Loup et 1 Peau de Troll)\n")
	Defilerapide("3. Bottes de l'aventurier (1 Fourrure de Loup et 1 Cuir de Sanglier)\n")
	Defilerapide("0. Quitter\n")
	var choixforgeron int
	var prix int
	fmt.Scan(&choixforgeron)

	switch choixforgeron {
	case 1:
		prix = 5
		p.Nomoney(prix)
		if p.Scan("Plume de Corbeau") && p.Scan("Cuir de Sanglier") {
			p.RemoveInventory("Plume de Corbeau")
			p.RemoveInventory("Cuir de Sanglier")
			p.Addinventory("Chapeau de l'aventurier")
		} else {
			Defilerapide("Vous n'avez pas les matériaux nécessaires\n")
		}

	case 2:
		prix = 5
		p.Nomoney(prix)
		if p.Scan("Peau de Troll") && p.Scan("Fourrure de Loup") {
			p.RemoveInventory("Fourrure de Loup")
			if p.Scan("Fourrure de Loup") {
				p.RemoveInventory("Fourrure de Loup")
				p.RemoveInventory("Peau de Troll")
				p.Addinventory("Tunique de l'aventurier")
			} else {
				Defilerapide("Vous n'avez pas assez de matériaux, Je vous rend votre Fourrure de Loup")
				p.Addinventory("Fourrure de Loup")
			}
		} else {
			Defilerapide("Vous n'avez pas les matériaux nécessaires\n")
		}
	case 3:
		prix = 5
		p.Nomoney(prix)
		if p.Scan("Fourrure de Loup") && p.Scan("Cuir de Sanglier") {
			p.RemoveInventory("Fourrure de Loup")
			p.RemoveInventory("Cuir de Sanglier")
			p.Addinventory("Bottes de l'aventurier")
		} else {
			Defilerapide("Vous n'avez pas les matériaux nécessaires")
		}
	case 0:
		p.Affichermenu()

	}
}
