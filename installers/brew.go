package installers

type Brew struct {
	InstallerBase
	Cask bool
}

// func (b *Brew)

func (b *Brew) test() {
	b.Name = "hi"
}
