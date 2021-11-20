package main

import (
	"tour_of_go/basic_types"
	"tour_of_go/constants"
	"tour_of_go/functions"
	"tour_of_go/imports"
	"tour_of_go/packages"
	"tour_of_go/variables"
)

func main() {
	packages.Main()
	imports.Main()
	functions.Main()
	variables.Main()
	basic_types.Main()
	constants.Main()
}
