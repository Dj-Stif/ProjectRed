package main

import "fmt"

// Création structure Personnage
type personnage struct {
	nom        string
	classe     string
	niveau     int
	PV_max     int
	PV_actuel  int
	inventaire []string
}

// Fonction pour initialiser personnage
func (p *personnage) init(nom string, classe string, niveau int, PV_max int, PV_actuel int, inventaire []string) {
	p.nom = nom
	p.classe = classe
	p.niveau = niveau
	p.PV_max = PV_max
	p.PV_actuel = PV_actuel
	p.inventaire = inventaire
}

// Fonction pour afficher info persos
func (p *personnage) displayinfo() {
	var quittermenu int
	fmt.Println("Information du personnage : ")
	fmt.Println("Nom :", p.nom)
	fmt.Println("classe :", p.classe)
	fmt.Println("niveau :", p.niveau)
	fmt.Println("PV Max :", p.PV_max)
	fmt.Println("PV_actuel :", p.PV_actuel)
	fmt.Println("Inventaire : ", p.inventaire)
	fmt.Println("--------------------------------")
	fmt.Println("0. Quitter")
	fmt.Scan(&quittermenu)

	switch quittermenu {
	case 0:
		p.affichermenu()
	default:
		fmt.Println("Choix invalide, veuillez réessayer")
	}
}

// Fonction Accès à l'inventaire

func (p *personnage) accessinventory() {
	var quitterinventaire int
	var choixinventaire int
	if len(p.inventaire) == 0 {
		fmt.Println("Votre inventaire est vide")
	} else {
		fmt.Println("Votre inventaire : ")
		for i, item := range p.inventaire {
			fmt.Printf("%d : %s\n", i+1, item)

			fmt.Scan(&choixinventaire)

			switch choixinventaire {
			case i + 1:
				if item == "Potion" {

				}
			}
		}
		fmt.Println("0. Quitter")
		fmt.Scan(&quitterinventaire)
		switch quitterinventaire {
		case 1:
			p.affichermenu()
		default:
			fmt.Println("Choix invalide, veuillez réessayer")
		}
	}

}
