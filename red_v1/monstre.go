package red

import "fmt"

//Structure Monstre
type Monster struct {
	monstername       string
	monsterPV_max     int
	monsterPV_actu    int
	monsterAttaque    int
	monsterinitiative int
}

// Fonction pour initialiser les monstres
func (m *Monster) Initgob(Mnom string, MPV_max int, MPV_actu int, Mattaque int, Minitiative int) {

	m.monstername = Mnom
	m.monsterPV_max = MPV_max
	m.monsterPV_actu = MPV_actu
	m.monsterAttaque = Mattaque
	m.monsterinitiative = Minitiative
}

// INIT MONSTER
var M1 Monster

// Fonction DisplayInfo Monstre
func (m *Monster) Displayinfomonster() {

	var quittermenumonstre int
	fmt.Println("Information du personnage : ")
	fmt.Println("Nom :", m.monstername)
	fmt.Println("PV Max :", m.monsterPV_max)
	fmt.Println("PV actuel :", m.monsterPV_actu)
	fmt.Println("Attaque :", m.monsterAttaque)
	fmt.Println("--------------------------------")
	fmt.Println("0. Quitter")
	fmt.Scan(&quittermenumonstre)

	switch quittermenumonstre {
	case 0:
		p.Affichermenu()
	default:
		fmt.Println("Choix invalide, veuillez réessayer")
	}
}

//Fonction pattern du gobelin
func (m *Monster) GoblinPattern(p *Personnage) {

	var degatsmonstre int

	degatsmonstre = m.monsterAttaque

	p.PV_actuel -= degatsmonstre

	fmt.Printf("%s à infligé %d dégats à %s \n", m.monstername, degatsmonstre, p.nom)
	p.Dead()
}
