package red

import "fmt"

func (p *Personnage) Addskillbook(selectedItem string) {
	if p.Scanskill("Livre de sort : Boule de feu [20 mana]") == false {
		p.skillbook = append(p.skillbook, selectedItem)
		fmt.Printf("Ajout de %s au skillbook \n", selectedItem)
		p.RemoveInventory(selectedItem)
	} else {
		fmt.Printf("Vous avez déjà ajouté %s \n", selectedItem)

	}
}

func (p *Personnage) Skill(M *Monster) {
	var choixskill int
	var degatsskill int
	fmt.Println("Vos skill : ")
	for i, item := range p.skillbook {
		fmt.Printf("%d : %s\n", i+1, item)
	}
	fmt.Println("0. Quitter")
	fmt.Scan(&choixskill)

	switch choixskill {
	case 0:
		p.Charturn(M)
	default:
		selectedItem := p.skillbook[choixskill-1]
		fmt.Printf("Vous avez sélectionné : %s\n", selectedItem)
		if selectedItem == "Livre de sort : Boule de feu [20 mana]" {
			if p.mana-20 < 0 {
				fmt.Println("Vous n'avez pas assez de mana")
				p.Skill(&Monster{})
			} else {
				p.mana -= 20
				degatsskill = 12
				M.monsterPV_actu -= degatsskill
				if M.monsterPV_actu < 0 {
					M.monsterPV_actu = 0
				}
				fmt.Printf("%s à infligé %d dégats à %s \n", p.nom, degatsskill, M.monstername)

			}
		}
		if selectedItem == "Coup de Poing [12 mana]" {
			if p.mana-12 < 0 {
				fmt.Println("Vous n'avez pas assez de mana")
			} else {
				p.mana -= 12
				degatsskill = 8
				M.monsterPV_actu -= degatsskill
				fmt.Printf("%s à infligé %d dégats à %s \n", p.nom, degatsskill, M.monstername)
			}
		}
	}
}
