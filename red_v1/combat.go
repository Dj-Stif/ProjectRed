package red

import (
	"fmt"
	"math/rand"
)

// Fonction CHAR TURN
func (p *Personnage) Charturn(m *Monster) {

	var choixcombat int
	Defilerapide("---------------------------------------Menu-------------------------------\n")
	Defilerapide("1. Attaque\n")
	Defilerapide("2. Inventaire\n")
	Defilerapide("3. Skill\n")
	Defilerapide("0. Quitter\n")
	fmt.Scan(&choixcombat)

	switch choixcombat {

	case 1:
		m.monsterPV_actu -= 5
		if m.monsterPV_actu < 0 {
			m.monsterPV_actu = 0
		}
		m.Monstermort(p)
		fmt.Printf("%s à infligé 5 points de  dégats à %s \n", p.nom, m.monstername)
	case 2:
		p.Accessinventorycombat()
	case 3:
		p.Skill(&M1)
		M1.Monstermort(p)
	case 0:
		p.Affichermenu()
	default:
		fmt.Println("Mauvais numéro veuillez réessayer")
	}

}

// ACCES INVENTAIRE VIA COMBAT
func (p *Personnage) Accessinventorycombat() {
	var choixinventairecombat int
	if len(p.inventaire) == 0 {
		fmt.Println("Votre inventaire est vide")
	} else {
		fmt.Println("Votre inventaire : ")
		for i, item := range p.inventaire {
			fmt.Printf("%d : %s\n", i+1, item)
		}

	}
	fmt.Println("0. Quitter")
	fmt.Scan(&choixinventairecombat)

	switch {
	case choixinventairecombat > len(p.inventaire):
		fmt.Println("Choix invalide. Veuillez sélectionner un numéro valide.")
		p.Accessinventorycombat()
	case choixinventairecombat == 0:
		p.Charturn(&Monster{})
	default:
		selectedItem := p.inventaire[choixinventairecombat-1]
		fmt.Printf("Vous avez sélectionné : %s\n", selectedItem)
		if selectedItem == "Potion" {
			p.Takepotcombat()
		} else {
			fmt.Println("Item non utilisable !")
			p.Accessinventorycombat()
		}
	}
}

// FONCTION COMBAT
func (p *Personnage) Combat(M1 *Monster) {
	nbtour := 1
	for p.PV_actuel >= 0 || M1.monsterPV_actu >= 0 {
		fmt.Printf("Vous êtes au %d ème tour \n", nbtour)
		if p.initiative < M1.monsterinitiative && nbtour == 1 {
			fmt.Printf("Le %s est plus rapide que vous !\n", M1.monstername)
			M1.GoblinPattern(p)
			p.Charturn(M1)
		} else if p.initiative == M1.monsterinitiative && nbtour == 1 {
			randomChoice := rand.Intn(2)
			if randomChoice == 0 {
				p.Charturn(M1)
				M1.GoblinPattern(p)
			}
			if randomChoice == 1 {
				fmt.Printf("Le %s est plus rapide que vous !\n", M1.monstername)
				M1.GoblinPattern(p)
				p.Charturn(M1)

			}
		} else {
			p.Charturn(M1)
			M1.GoblinPattern(p)
		}

		fmt.Printf("%s à : %d / %d PV \n", p.nom, p.PV_actuel, p.PV_max)
		fmt.Printf("%s à : %d / %d mana \n", p.nom, p.mana, p.manamax)
		fmt.Printf("%s à %d / %d PV \n", M1.monstername, M1.monsterPV_actu, M1.monsterPV_max)
		nbtour++
	}

}
