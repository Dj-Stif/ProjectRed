package red

import "fmt"

func (p *Personnage) Qui() {
	var choixqui int
	fmt.Println("-------------------------Qui sont-ils ?----------------------------")
	Defilerapide("Partie 2 : ABBA\n")
	Defilerapide("Partie 3 : Steven Spielberg\n")
	Defilerapide("Partie 4 : Queen\n")
	Defilerapide("0. Quitter\n")
	fmt.Scan(&choixqui)
	switch choixqui {
	case 0:
		p.Affichermenu()
	}
}
