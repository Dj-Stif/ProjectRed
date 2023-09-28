package red

import "fmt"

// Fonction Marchand

func (p *Personnage) Accessmarchand() {
	var choixmarchand int
	var prix int
	fmt.Printf("Vous avez %d Or\n", p.argent)
	Defilerapide("Bievnenue chez le marchand : \n")
	Defilerapide("1. Potion (0 or)\n")
	Defilerapide("2. Potion (3 or)\n")
	Defilerapide("3. Potion de poison (6 or)\n")
	Defilerapide("4. Livre de sort : Boule de feu (25 or)\n")
	Defilerapide("5. Fourrure de Loup (4 or)\n")
	Defilerapide("6. Peau de Troll (7 or)\n")
	Defilerapide("7. Cuir de Sanglier (3 or)\n")
	Defilerapide("8. Plume de Corbeau (1 or)\n")

	Defilerapide("0. Quitter le marchand.\n")
	fmt.Scan(&choixmarchand)

	switch choixmarchand {

	case 1:
		p.Addinventory("Potion")
		p.Accessmarchand()
	case 2:
		prix = 3
		if p.Nomoney(prix) {
			p.Addinventory("Potion")
			p.Accessmarchand()
		}

	case 3:
		prix = 6
		if p.Nomoney(prix) {
			p.Addinventory("Potion de poison")
		}
		p.Accessmarchand()

	case 4:
		prix = 25
		if p.Nomoney(prix) {
			p.Addinventory("Livre de sort : Boule de feu [20 mana]")
		}
		p.Accessmarchand()

	case 5:
		prix = 4
		if p.Nomoney(prix) {
			p.Addinventory("Fourrure de Loup")
		}
		p.Accessmarchand()

	case 6:
		prix = 7
		if p.Nomoney(prix) {
			p.Addinventory("Peau de Troll")
		}
		p.Accessmarchand()

	case 7:
		prix = 3
		if p.Nomoney(prix) {
			p.Addinventory("Cuir de Sanglier")
		}
		p.Accessmarchand()

	case 8:
		prix = 1
		if p.Nomoney(prix) {
			p.Addinventory("Plume de Corbeau")
		}
		p.Accessmarchand()
	case 0:
		p.Affichermenu()

	}
}
