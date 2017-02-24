// package packageInstaller
package main

import "fmt"

type Package struct {
	Name      string
	Version   string
	Installer string
}

func (p *Package) Install() {
	fmt.Println("Install")
}
