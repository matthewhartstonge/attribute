package attribute

import (
	// Standard Library Imports
	"fmt"

	// External Imports
	"github.com/sirupsen/logrus"
)

func GenerateAttributions(config *Config) []byte {
	attribution := config.Project.AttributionType
	template, ok := attributions[attributionType(attribution)]
	if !ok {
		errMsg := fmt.Sprintf("Unable to generate attibutions for attribution type '%s'", attribution)
		logrus.Fatal(errMsg)
	}

	attributionBytes := GenerateFromTemplate(attribution, template, config)
	return attributionBytes
}
