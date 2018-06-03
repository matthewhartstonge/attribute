package attribute

// Project contains meta information to configure the licenses and Open Source
// Attributions to be generated.
type Project struct {
	// Year specifies the year to be entered into your copyright notice
	Year string
	// Owner specifies the owner of the project
	Owner string
	// Custom specifies an email address to contact the project owner on, or
	// anything extra you want added to your copyright notice as it's appended
	// to the copyright string.
	Custom string

	// LicenseType is the type of license you want to use for the project.
	LicenseType string `yaml:"licenseType"`
	// AttributionType specifies the type of Open Source Attribution you
	// want to generate.
	AttributionType string `yaml:"attributionType"`
}

// Dependency contains meta information about a dependency used within the
// project.
type Dependency struct {
	// Name is the dependency project name
	Name string
	// Link contains the link to the repository.
	Link string
	// Copyright contains the extracted copyright string from the License or
	// Notice file.
	Copyright string
	// License contains the full text of the license.
	License string
	// LicenseLink contains a direct link to the license within the dependencies
	// repository.
	LicenseLink string `yaml:"licenseLink"`
	// LicenseName contains the string form of the license that the project uses.
	LicenseName string `yaml:"licenseName"`
	// licenseFile is used internally to generate links based on file name.
	licenseFile string
}

// Config contains the data structure to unmarshal the .attribute.yaml file.
type Config struct {
	Project      Project      `yaml:"project"`
	Attributions []Dependency `yaml:"attributions"`
}
