package red

import (
	"fmt"
	"time"
)

func (p *Personnage) Afficherdemarrage() {
	var choixdemarrage int

	for {
		Defile("Bienvenue dans notre jeu !\nAppuyez sur 1 pour démarrer !\n")

		fmt.Scan(&choixdemarrage)

		switch choixdemarrage {
		case 1:
			p.CharCreation()
		default:
			Defile("Choix invalide, veuillez réessayer\n")
		}
	}
}

func (p *Personnage) Affichermenu() {

	var choixmenu int
	for {
		fmt.Println("---------------------------------MENU------------------------------- :")
		Defilerapide("1. Afficher les informations du personnage\n")
		Defilerapide("2. Accèder au contenu de l'inventaire\n")
		Defilerapide("3. Marchand\n")
		Defilerapide("4. Forgeron\n")
		Defilerapide("5. Menucombat\n")
		Defilerapide("6. Qui sont-ils ?\n")
		Defilerapide("0. Quitter\n")
		Defilerapide("Veuillez faire un choix :\n")
		fmt.Scan(&choixmenu)

		switch choixmenu {
		case 1:
			p.Displayinfo()
		case 2:
			p.Accessinventory()
		case 3:
			p.Accessmarchand()
		case 4:
			p.Accesforgeron()
		case 5:
			p.Combat(&M1)
		case 6:
			p.Qui()
		case 0:
			fmt.Println("Au revoir !")
			p.Afficherdemarrage()
		default:
			fmt.Println("Choix invalide, veuillez réessayer")

		}
	}
}
func Defile(s string) {
	for _, letter := range s {
		fmt.Print(string(letter))
		time.Sleep(30 * time.Millisecond)
	}
}

func Defilerapide(s string) {
	for _, letter := range s {
		fmt.Print(string(letter))
		time.Sleep(5 * time.Millisecond)
	}
}

func Defilepara(s string, p *Personnage) {
	for _, letter := range s {
		fmt.Print(string(letter))
		time.Sleep(5 * time.Millisecond)
	}
}
