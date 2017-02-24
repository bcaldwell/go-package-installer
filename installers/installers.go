package installers

import "fmt"

type Installer interface {
	Install()
}

type InstallerBase struct {
	Name    string
	Version string
}

func (b *InstallerBase) greet() {
	fmt.Println("hello: " + b.Name)
}
