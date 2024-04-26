package engine

type Structure struct {
	//variables en lien avec le web
	running            bool
	pagesWeblist       []string
	iteratorWebPage    int
	closeWindow		   bool

	//variables servant de changer l'état du jeu
	win                bool
	end                bool
	try                int

	listColorsCase     []string
	listVerifCase      []string
	codeColor		   []string
	listColors		   []string
	currentColor       string
	indiceColor        int
	indiceMinimum      int
	indiceMaximum      int
	indiceVerifCase    int

}

func (g *Structure) Run() {
	//lance le programme
	g.init()		
	if g.running {	 //lance en continue la fonction web
		g.web()
	}
}
