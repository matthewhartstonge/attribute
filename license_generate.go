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
		logrus.Error(fmt.Sprintf("Unable to generate license for license type '%s'", license))
		return nil
	}

	if template.Full == "" {
		return nil
	}

	return GenerateFromTemplate(license, template.Full, config)
}

func GenerateNotice(config *Config) []byte {
	license := config.Project.LicenseType
	template, ok := licenses[licenseType(license)]
	if !ok {
		logrus.Error(fmt.Sprintf("Unable to generate notice for license type '%s'", license))
		return nil
	}

	if template.Notice == "" {
		return nil
	}

	return GenerateFromTemplate(license, template.Notice, config)
}
