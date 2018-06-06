package attribute

type licenseType string

type license struct {
	// Type contains the license type for meta analysis
	Type licenseType

	// Pretty Name of the License for generating attributions
	Name string

	// Notice contains the license notice to be prepended to all Go files.
	Notice string

	// Full contains the license in full to be generated at the root of the
	// repo.
	Full string
}

// licenses binds all the licenses to generate together.
var licenses = map[licenseType]license{
	licenseTypeApache2: {
		Type:   licenseTypeApache2,
		Name:   licensePrettyNameApache2,
		Notice: licenseTemplateApache2Notice,
		Full:   licenseTemplateApache2Full,
	},
	licenseTypeBSD2Clause: {
		Type: licenseTypeBSD2Clause,
		Name: licensePrettyNameBSD2Clause,
		Full: licenseTemplateBSD2ClauseFull,
	},
	licenseTypeBSD3Clause: {
		Type: licenseTypeBSD3Clause,
		Name: licensePrettyNameBSD3Clause,
		Full: licenseTemplateBSD3ClauseFull,
	},
	licenseTypeMIT: {
		Type: licenseTypeMIT,
		Name: licensePrettyNameMIT,
		Full: licenseTemplateMITFull,
	},
}
