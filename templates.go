package comscape

import (
	"html/template"
)

// Tmpl is an exported variable
var Tmpl = template.Must(template.ParseGlob("../resources/templates/gohtml/*"))

type LandingPage struct {
	competitors []competitor
}

type Competitor struct {
	name        string
	description string
}
