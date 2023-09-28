package red

import (
	"fmt"
	"unicode"
)

var p Personnage

// Création structure Personnage
type Personnage struct {
	nom        string
	classe     string
	niveau     int
	PV_max     int
	PV_actuel  int
	inventaire []string
	argent     int
	equipement Equipement
	skillbook  []string
	mana       int
	manamax    int
	initiative int
}

// Structure Equipement
type Equipement struct {
	tête  string
	torse string
	pieds string
}

// Fonction pour initialiser personnage
func (p *Personnage) Init(choixnom string, classe string, niveau int, PV_max int, PV_actuel int, inventaire []string, argent int, tête string, torse string, pieds string, skillbook []string, mana int, manamax int, initiative int) {

	p.nom = choixnom
	p.classe = classe
	p.niveau = niveau
	p.PV_max = PV_max
	p.PV_actuel = PV_actuel
	p.inventaire = inventaire
	p.argent = argent
	p.equipement.tête = tête
	p.equipement.torse = torse
	p.equipement.pieds = pieds
	p.skillbook = skillbook
	p.mana = mana
	p.manamax = manamax
	p.initiative = initiative
}

// Fonction pour afficher info persos
func (p *Personnage) Displayinfo() {
	var quittermenu int
	Defilerapide("Information du personnage : ")
	fmt.Println("Nom :", p.nom)
	fmt.Println("classe :", p.classe)
	fmt.Println("niveau :", p.niveau)
	fmt.Println("PV Max :", p.PV_max)
	fmt.Println("PV_actuel :", p.PV_actuel)
	fmt.Println("Inventaire : ", p.inventaire)
	fmt.Println("Argent : ", p.argent)
	fmt.Println("Tête : ", p.equipement.tête)
	fmt.Println("Torse : ", p.equipement.torse)
	fmt.Println("Pieds : ", p.equipement.pieds)
	fmt.Println("Skillbook : ", p.skillbook)
	fmt.Println("Mana : ", p.mana)
	fmt.Println("Mana Max :", p.manamax)
	fmt.Println("Initiative :", p.initiative)
	fmt.Println("--------------------------------")
	fmt.Println("0. Quitter")
	fmt.Scan(&quittermenu)

	switch quittermenu {
	case 0:
		p.Affichermenu()
	default:
		Defilerapide("Choix invalide, veuillez réessayer")
	}
}

// Fonction Accès à l'inventaire

func (p *Personnage) Accessinventory() {

	var choixinventaire int
	if len(p.inventaire) == 0 {
		fmt.Println("Votre inventaire est vide")
	} else {
		fmt.Println("Votre inventaire : ")
		for i, item := range p.inventaire {
			fmt.Printf("%d : %s\n", i+1, item)
		}

	}
	fmt.Println("0. Quitter")
	fmt.Scan(&choixinventaire)

	switch {
	case choixinventaire > len(p.inventaire):
		fmt.Println("Choix invalide. Veuillez sélectionner un numéro valide.")
		p.Accessinventory()
	case choixinventaire == 0:
		p.Affichermenu()
	default:
		selectedItem := p.inventaire[choixinventaire-1]
		fmt.Printf("Vous avez sélectionné : %s\n", selectedItem)

		if selectedItem == "Livre de sort : Boule de feu [20 mana]" {
			p.Addskillbook(selectedItem)
			break
		}
		if selectedItem == "Potion" {
			p.Takepot()
			p.Accessinventory()
		}
		if selectedItem == "Chapeau de l'aventurier" {
			if p.equipement.tête != "Vide" {
				p.Addinventory(p.equipement.tête)
			}
			p.equipement.tête = "Chapeau de l'aventurier"
			fmt.Println("Vous avez équipé le Chapeau de l'aventurier")
			p.Upstats(10)
			p.RemoveInventory("Chapeau de l'aventurier")
			p.Accessinventory()
			break
		}
		if selectedItem == "Tunique de l'aventurier" {
			p.equipement.torse = "Tunique de l'aventurier"
			fmt.Println("Vous avez équipé la tunique de l'aventurier")
			p.Upstats(25)
			p.RemoveInventory("Tunique de l'aventurier")
			p.Accessinventory()
			break
		}
		if selectedItem == "Bottes de l'aventurier" {
			p.equipement.pieds = "Bottes de l'aventurier"
			p.Upstats(15)
			fmt.Println("Vous avez équipé les Bottes de l'aventurier")
			p.RemoveInventory("Bottes de l'aventurier")
			p.Accessinventory()
			break
		}
		if selectedItem == "Potion de poison" {
			p.Poisonpot()
			p.RemoveInventory(selectedItem)
		} else {
			fmt.Println("Item non utilisable !")
			p.Accessinventory()
		}
	}
}

// Fonction UP de stats :
func (p *Personnage) Upstats(PV_bonus int) {

	if p.equipement.tête == "Chapeau de l'aventurier" {
		p.PV_max += PV_bonus
	}
	if p.equipement.tête == "Tunique de l'aventurier" {
		p.PV_max += PV_bonus
	}
	if p.equipement.tête == "Bottes de l'aventurier" {
		p.PV_max += PV_bonus
	}
}

// Fonction Char creation
func (p *Personnage) CharCreation() {
	var choixnom string
	for {
		Defile("Entrez votre nom : ")
		fmt.Scan(&choixnom)
		if p.Nomvalide(choixnom) {
			break
		} else {
			Defilerapide("Le nom ne doit contenir que des lettres")
		}
	}
	for {
		var choixclasse int
		Defilerapide("Choissisez votre classe : \n")
		Defilerapide("1. Humain, moyen dans tout les domaines\n")
		Defilerapide("2. Elfe, plus fragile que les autres mais aussi plus agile \n")
		Defilerapide("3. Nain, petit lent mais robuste \n")
		fmt.Scan(&choixclasse)

		switch choixclasse {
		case 1:
			p.Init(choixnom, "Humain", 1, 100, 50, []string{}, 100, "Vide", "Vide", "Vide", []string{"Coup de Poing [12 mana]"}, 100, 100, 3)
		case 2:
			p.Init(choixnom, "Elfe", 1, 80, 40, []string{}, 100, "Vide", "Vide", "Vide", []string{"Coup de Poing [12 mana]"}, 100, 100, 5)
		case 3:
			p.Init(choixnom, "Nain", 1, 120, 60, []string{}, 100, "Vide", "Vide", "Vide", []string{"Coup de Poing [12 mana]"}, 100, 100, 1)
		default:
			fmt.Println("Vous n'avez pas choisi de classe")
			continue
		}

		p.Affichermenu()
	}

}

// NOM VALIDE OU PAS
func (p *Personnage) Nomvalide(nom string) bool {
	for _, lettres := range nom {
		if !unicode.IsLetter(lettres) {
			return false
		}
	}
	return true
}
