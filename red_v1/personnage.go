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
	argent     int
}

// Fonction pour initialiser personnage
func (p *personnage) init(nom string, classe string, niveau int, PV_max int, PV_actuel int, inventaire []string, argent int) {
	p.nom = nom
	p.classe = classe
	p.niveau = niveau
	p.PV_max = PV_max
	p.PV_actuel = PV_actuel
	p.inventaire = inventaire
	p.argent = argent
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
	fmt.Println("Argent : ", p.argent)
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

	var choixinventaire int
	if len(p.inventaire) == 0 {
		fmt.Println("Votre inventaire est vide")
	} else {
		fmt.Println("Votre inventaire : ")
		for i, item := range p.inventaire {
			fmt.Printf("%d : %s\n", i+1, item)
		}
		fmt.Println("0. Quitter")
	}
	fmt.Scan(&choixinventaire)
	switch choixinventaire {
	case len(p.inventaire):
		fmt.Println("Choix invalide. Veuillez sélectionner un numéro valide.")
	case 0:
		p.affichermenu()
	default:
		selectedItem := p.inventaire[choixinventaire-1]
		fmt.Printf("Vous avez sélectionné : %s\n", selectedItem)
		if selectedItem == "Potion" {
			p.takepot()
			p.accessinventory()
		}
	}
}
