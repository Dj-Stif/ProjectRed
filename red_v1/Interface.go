// Fonction affichage menu
package red

import "fmt"

func (p *personnage) afficherdemarrage() {
	var choixdemarrage int
	for {
		fmt.Println("1. Démarrer")
		fmt.Println("2. Paramètres")
		fmt.Println("3. Quitter")
		fmt.Scan(&choixdemarrage)

		switch choixdemarrage {
		case 1:
			p.choixpersonnage()
		case 2:
			fmt.Println("En cours...")
		case 3:
			fmt.Println("Au revoir !")
		default:
			fmt.Println("Choix invalide, veuillez réessayer")

		}
	}
}

func (p *personnage) choixpersonnage() {
	var choixperso int
	fmt.Println("Select your character : ")
	fmt.Println("1. Elfe")
	fmt.Scan(&choixperso)

	switch choixperso {
	case 1:
		var p1 personnage
		p1.init("Marin", "elfe", 1, 100, 40, []string{"Potion", "Potion", "Potion"})
		p1.affichermenu()

	}

}

func (p *personnage) affichermenu() {

	var choixmenu int
	for {
		fmt.Println("MENU :")
		fmt.Println("1. Afficher les informations du personnage ")
		fmt.Println("2. Accèder au contenu de l'inventaire ")
		fmt.Println("3. Marchand ")
		fmt.Println("4. Quitter")
		fmt.Println("Veuillez faire un choix : ")
		fmt.Scan(&choixmenu)

		switch choixmenu {
		case 1:
			p.displayinfo()
		case 2:
			p.accessinventory()
		case 3:
			p.accessmarchand()
		case 4:
			fmt.Println("Au revoir !")
			p.afficherdemarrage()
		default:
			fmt.Println("Choix invalide, veuillez réessayer")

		}
	}
}
