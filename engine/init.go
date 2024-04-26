package engine

import (
	"fmt"
	"math/rand"
	"time"
)

func (g *Structure) init() {
	// initialisation de toutes les variables de la structure Engine

	//initialise les variables servant de changer l'état du jeu
	g.try = 0
	g.win = false
	g.end = false

	// initialisation des variables en lien avec le web
	g.pagesWeblist = []string{"templates/index.html", "templates/game.html"}
	g.iteratorWebPage = 0
	g.running = true
	g.closeWindow = false

	g.currentColor = ""
	g.listColorsCase = []string{}
	g.listVerifCase = []string{}
	g.listColors = []string{"bleu", "rouge", "jaune", "vert", "violet", "orange"}
	g.codeColor = []string{}
	g.indiceColor = 0
	g.indiceMinimum = 0
	g.indiceMaximum = 3
	g.indiceVerifCase = 0

	//initialisation des variables "randoms"

	for i := 0; i < 44; i++ {
		g.listColorsCase = append(g.listColorsCase, "blanc")
		g.listVerifCase = append(g.listVerifCase, "blanc")
	}

	s := rand.NewSource(time.Now().Unix())
	for i := 0; i < 4; i++ {
		r := rand.New(s)
		g.codeColor = append(g.codeColor, g.listColors[r.Intn(6)])
	}
	fmt.Println(g.codeColor)
}
