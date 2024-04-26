package engine

import (
	"fmt"
	"html/template"
	"net/http"
)

// initialisation de la structure DrawWeb
type DrawWeb struct {
	Color string
	Colors []string
	ColorsTry []string
	Win bool
	Try int
	End bool
}


func(g *Structure) web() {
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	http.Handle("/pictures/", http.StripPrefix("/pictures/", http.FileServer(http.Dir("pictures"))))
	//http.Handle("/javascript/", http.StripPrefix("/javascript/", http.FileServer(http.Dir("javascript"))))
	//http.Handle("/police/", http.StripPrefix("/police/", http.FileServer(http.Dir("police"))))
	//http.Handle("/audio/", http.StripPrefix("/audio/", http.FileServer(http.Dir("audio"))))
	// chargement de toutes les pages"
	http.HandleFunc("/", g.index)
	http.HandleFunc("/index", g.index)
	http.HandleFunc("/game", g.game)

	// chargement du port utilisé
	fmt.Println("http://localhost:8090/")
	http.ListenAndServe(":8090", nil)
	
}

func(g *Structure) index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(g.pagesWeblist[0]))
	tmpl.Execute(w, nil)
}

func(g *Structure) game(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  
	tmpl := template.Must(template.ParseFiles(g.pagesWeblist[1]))
	g.currentColor = r.Form.Get("color")
	if len(g.currentColor) > 0 {
		if g.currentColor == "delete" {
			if g.indiceColor > g.indiceMinimum {
				g.listColorsCase[g.indiceColor-1] = "blanc"
				g.indiceColor--
			}
		} else if g.currentColor == "valide" {
			if g.indiceColor-1 == g.indiceMaximum {
				g.verifLine();
				g.indiceMaximum += 4
				g.indiceMinimum += 4
			}				
		} else if g.indiceColor < 44 && g.indiceColor >= g.indiceMinimum && g.indiceColor <= g.indiceMaximum {
			g.listColorsCase[g.indiceColor] = g.currentColor
			g.currentColor = ""
			g.indiceColor++
		}
	}
	
	web := DrawWeb {
		Color: g.currentColor,
		Colors: g.listColorsCase,
		ColorsTry: g.listVerifCase,
		Win: g.win,
		Try: g.try,
		End: g.end,
	}
	tmpl.Execute(w, web)  // permet d'envoyer les valeurs de la structure DrawWeb sur le html

	if g.closeWindow {
		g.init()
		return
	}
}

