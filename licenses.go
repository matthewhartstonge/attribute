package attribute

import (
	// Standard Library Imports
	"regexp"
)

type licenseType string

type license struct {
	// Stub contains the license stub to be prepended to all Go files.
	Stub string

	// Full contains the license in full to be generated at the root of the
	// repo.
	Full string

	// CopyRightRegex enables searching an incoming license to find related
	// copyright info which we will generate attribution files from.
	CopyRightRegex regexp.Regexp
}

// licenses binds all the licenses to generate together.
var licenses = map[licenseType]license{
	licenseApache2: license{
		Stub: licenseTemplateApache2Stub,
		Full: licenseTemplateApache2Full,
	},
}
