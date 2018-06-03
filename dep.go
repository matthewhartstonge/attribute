package attribute

import (
	// Standard Library Imports
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	// External Imports
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

// expectedSimilarity provides the set probability that the template license matches
// the license provided in the repository.
const expectedSimilarity = 0.9
const fileDep = "Gopkg.toml"
const fileNotice = "NOTICE"
const fileLicense = "LICENSE"

var versionRegex *regexp.Regexp

// toml structs pulled from dep
type rawManifest struct {
	Constraints  []rawProject    `toml:"constraint,omitempty"`
	Overrides    []rawProject    `toml:"override,omitempty"`
	Ignored      []string        `toml:"ignored,omitempty"`
	Required     []string        `toml:"required,omitempty"`
	PruneOptions rawPruneOptions `toml:"prune,omitempty"`
}

type rawProject struct {
	Name     string `toml:"name"`
	Branch   string `toml:"branch,omitempty"`
	Revision string `toml:"revision,omitempty"`
	Version  string `toml:"version,omitempty"`
	Source   string `toml:"source,omitempty"`
}

type rawPruneOptions struct {
	UnusedPackages bool `toml:"unused-packages,omitempty"`
	NonGoFiles     bool `toml:"non-go,omitempty"`
	GoTests        bool `toml:"go-tests,omitempty"`

	//Projects []map[string]interface{} `toml:"project,omitempty"`
	Projects []map[string]interface{}
}

func init() {
	regex, _ := regexp.Compile(`[~^]`)
	versionRegex = regex
}

func GetDependencies() []*Dependency {
	logger := logrus.WithFields(logrus.Fields{
		"package": "attribute",
		"method":  "getConstraints",
	})

	f, err := ioutil.ReadFile(fileDep)
	if err != nil {
		logger.WithError(err).Fatal("Unable to read Gopkg.toml file")
	}

	manifest := &rawManifest{}
	err = toml.Unmarshal(f, manifest)
	if err != nil {
		logger.WithError(err).Fatal("Unable to parse Gopkg.toml file")
	}

	var deps []*Dependency
	constraints := manifest.Constraints
	for _, constraint := range constraints {
		dep := &Dependency{}
		SetDependencyName(constraint.Name, dep)
		SetDependencyLicenseMeta(constraint.Name, dep)
		SetDependencyLinks(constraint, dep)

		deps = append(deps, dep)
	}

	return deps
}

func SetDependencyName(depPath string, dependency *Dependency) *Dependency {
	splitDepPath := strings.Split(depPath, "/")
	dependency.Name = splitDepPath[len(splitDepPath)-1]

	return dependency
}

// setDependencyLinks builds up links.
// For a link to a project's license, it requires the dependency meta to have been set..
func SetDependencyLinks(constraint rawProject, dependency *Dependency) *Dependency {
	var baselink = constraint.Name
	if constraint.Source != "" {
		baselink = constraint.Source
	}
	dependency.Link = fmt.Sprintf("https://%s", baselink)

	if dependency.licenseFile != "" {
		var sublink string
		if constraint.Branch != "" {
			sublink = constraint.Branch
		}
		if constraint.Revision != "" {
			sublink = constraint.Revision
		}
		if constraint.Version != "" {
			// strip dep version
			sublink = constraint.Version
			if versionRegex.MatchString(constraint.Version) {
				sublink = constraint.Version[1:len(constraint.Version)]
			}

			// Convention dictates prefixing version numbers with v
			firstChar := string(sublink[0])
			if _, err := strconv.Atoi(firstChar); err == nil {
				sublink = "v" + sublink
			}
		}
		dependency.LicenseLink = fmt.Sprintf("%s/blob/%s/%s", dependency.Link, sublink, dependency.licenseFile)
	}
	return dependency
}

func SetDependencyLicenseMeta(depPath string, dependency *Dependency) *Dependency {
	logger := logrus.WithFields(logrus.Fields{
		"package": "attribute",
		"method":  "getLicense",
	})

	var licenseTemplate *license

	// Prefer a notice to a license file..
	dependency.licenseFile = fileNotice
	filePath := fmt.Sprintf("./vendor/%s/%s", depPath, fileNotice)
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		errmsg := fmt.Sprintf("unable to find/read notice file for '%s'", depPath)
		logger.WithError(err).Debug(errmsg)

		dependency.licenseFile = fileLicense
		filePath = fmt.Sprintf("./vendor/%s/%s", depPath, fileLicense)
		file, err = ioutil.ReadFile(filePath)
		if err != nil {
			errmsg := fmt.Sprintf("unable to read license file for '%s'", depPath)
			logger.WithError(err).Error(errmsg)
			return dependency
		}
	}

	vendorLicense := strings.Split(string(file), "\n")
	for licType, licenseMeta := range licenses {
		var tmpl string
		switch dependency.licenseFile {
		case fileNotice:
			tmpl = licenseMeta.Notice
		case fileLicense:
			tmpl = licenseMeta.Full
		}

		similarityRatio := levenshtein.RatioForStrings([]rune(tmpl), []rune(string(file)), levenshtein.DefaultOptions)
		if similarityRatio >= expectedSimilarity {
			logger.WithField("similarityRatio", similarityRatio).
				Debug(fmt.Sprintf("Found License for %s, type appears to be %s", depPath, licType))
			licenseTemplate = &licenseMeta
			break
		} else {
			logger.WithField("similarityRatio", similarityRatio).
				Debug(fmt.Sprintf("Found License for %s, does not appear to be an %s based license.", depPath, licType))
		}
	}

	// unknown license
	if licenseTemplate == nil {
		return nil
	}

	//// process license meta
	// extract copyright comment
	var found string
	for _, line := range vendorLicense {
		searchText := strings.TrimSpace(line)
		if strings.HasPrefix(strings.ToLower(searchText), "copyright ") {
			found = line
		}
		if found != "" {
			logger.Debug(fmt.Sprintf("Found the copyright string in %s's license file", depPath))
			break
		}
	}

	dependency.Copyright = found
	dependency.License = string(file)
	dependency.LicenseName = string(licenseTemplate.Name)
	return dependency
}
