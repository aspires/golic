package templates

type License struct {
	Name     string
	URL      string
	Template string
}

func List() []string {
	licenses := []string{}

	for i := range Licenses {
		licenses = append(licenses, Licenses[i].Name)
	}

	return licenses
}

func Load(name string) (*License, bool) {
	for i := range Licenses {
		if Licenses[i].Name == name {
			return &Licenses[i], true
		}
	}

	return nil, false
}
