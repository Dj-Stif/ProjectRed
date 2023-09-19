package piscine

type Personnage struct {
	race       string
	classe     string
	niveau     int
	PvMax      int
	Pv         int
	inventaire map[string]int
	argent     int
}

func (p Personnage) Init(r string, c string, n int, pvmax int, pv int, argent int) {
	p.race = r
	p.classe = c
	p.niveau = n
	p.PvMax = pvmax
	p.Pv = pv
	p.inventaire = map[string]int{"potion de vie": 1}
	p.argent = argent
}
