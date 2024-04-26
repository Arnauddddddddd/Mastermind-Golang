package engine

func (g *Structure) verifLine() {
	g.indiceVerifCase = g.indiceMinimum
	g.indiceVerifCase = g.indiceMinimum
	codeCopy := []string{}
	caseCopy := []string{}
	for i := 0; i < len(g.codeColor); i++ {
		codeCopy = append(codeCopy, g.codeColor[i])
	}
	for i := 0; i < len(g.listColorsCase); i++ {
		caseCopy = append(caseCopy, g.listColorsCase[i])
	}

	for i := 0; i < 4; i++ {
		if caseCopy[i+g.indiceMinimum] == codeCopy[i] && codeCopy[i] != "x" && caseCopy[i+g.indiceMinimum] != "x"{
			codeCopy[i] = "x"
			caseCopy[i+g.indiceMinimum] = "x"
			g.listVerifCase[g.indiceVerifCase] = "rouge"
			g.indiceVerifCase++
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if caseCopy[i+g.indiceMinimum] == codeCopy[j] && codeCopy[j] != "x" && caseCopy[i+g.indiceMinimum] != "x"{
				codeCopy[j] = "x"
				caseCopy[i+g.indiceMinimum] = "x"
				g.listVerifCase[g.indiceVerifCase] = "gris"
				g.indiceVerifCase++
			}
		}		
	}

}
