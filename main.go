// Execution
package main

import red "red/red_v1"

func main() {
	var p red.Personnage
	red.M1.Initgob("Gobelin d'entrainement", 40, 40, 5, 3)
	p.Afficherdemarrage()
}
