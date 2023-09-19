package red

import "fmt"

// Fonction Marchand
func (p *Personnage) Accessmarchand() {
	fmt.Println("Bievenue chez le marchand : ")
	fmt.Println("1. Potion (0 or)")
	fmt.Println("Quitter le marchand.")
}
