package attribute

import (
	// Standard Library Imports
	"bytes"
	"fmt"
	"text/template"

	// External Library Imports
	"github.com/sirupsen/logrus"
)

func GenerateFromTemplate(templateName, templateContents string, config *Config) []byte {
	tmpl, err := template.New(templateName).Parse(templateContents)
	if err != nil {
		logrus.WithError(err).Fatal(fmt.Sprintf("Unable to load template for %s successfully", templateName))
	}

	var generatedLicense bytes.Buffer
	err = tmpl.Execute(&generatedLicense, config)
	if err != nil {
		logrus.WithError(err).Fatal(fmt.Sprintf("Unable to generate %s successfully", templateName))
	}

	return generatedLicense.Bytes()
}
