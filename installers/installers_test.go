package installers

import "testing"

func TestInstallerBase_greet(t *testing.T) {
	type fields struct {
		Name    string
		Version string
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &InstallerBase{
				Name:    tt.fields.Name,
				Version: tt.fields.Version,
			}
			b.greet()
		})
	}
}

func TestInstallerBase_greet_2(t *testing.T) {
	a := &Brew{
		InstallerBase: InstallerBase{
			Name: "boo",
		},
	}
	a.test()
	a.greet()
}
