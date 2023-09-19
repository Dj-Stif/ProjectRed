package main

import "fmt"

// Fonction Marchand
func (p *personnage) accessmarchand() {
	var choixmarchand int
	fmt.Println("Bievenue chez le marchand : ")
	fmt.Println("1. Potion (0 or)")
	fmt.Println("2. Potion (3 or)")
	fmt.Println("3. Potion de poison (6 or)")
	fmt.Println("4. Livre de sort : Boule de feu (25 or)")
	fmt.Println("5. Fourrure de Loup (4 or)")
	fmt.Println("6. Peau de Troll (7 or)")
	fmt.Println("7. Cuir de sanglier (3 or)")
	fmt.Println("8. Plume de corbeau (1 or)")

	fmt.Println("0. Quitter le marchand.")
	fmt.Scan(&choixmarchand)

	switch choixmarchand {

	case 1:
		p.addinventory("Potion")
		p.accessmarchand()
	case 2:
		p.addinventory("Potion")
		p.argent -= 3
		p.accessmarchand()

	case 3:
		p.addinventory("Potion de poison")
		p.argent -= 6
		p.accessmarchand()

	case 4:
		p.addinventory("Livre de sort : Boule de feu")
		p.argent -= 25
		p.accessmarchand()

	case 5:
		p.addinventory("Fourrure de Loup")
		p.argent -= 4
		p.accessmarchand()

	case 6:
		p.addinventory("Peau de Troll")
		p.argent -= 7
		p.accessmarchand()

	case 7:
		p.addinventory("Cuir de sanglier")
		p.argent -= 3
		p.accessmarchand()

	case 8:
		p.addinventory("Plume de Corbeau")
		p.argent -= 1
		p.accessmarchand()

	case 0:
		p.affichermenu()

	}
}
