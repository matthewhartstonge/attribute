package attribute

import (
	// Standard Library Imports
	"fmt"

	// External Imports
	"github.com/sirupsen/logrus"
)

func GenerateLicense(config *Config) []byte {
	license := config.Project.LicenseType
	template, ok := licenses[licenseType(license)]
	if !ok {
		logrus.Fatal(fmt.Sprintf("Unable to generate license for license type '%s'", license))
	}

	licenseBytes := GenerateFromTemplate(license, template.Full, config)
	return licenseBytes
}
