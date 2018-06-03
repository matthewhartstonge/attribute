package attribute

const attributionTypeShort = "short"
const attributionShort = `# Open Source License Attribution

This application uses Open Source components. You can find the source
code of their open source projects along with license information below.
We acknowledge and are grateful to these developers for their contributions
to open source.
{{range .Attributions}}
{{if .Link}}### [{{.Name}}]({{.Link}}){{else}}### {{.Name}}{{end}}
{{if .Copyright}}- {{.Copyright}}{{- else -}}{{- end}}
{{if .LicenseName}}{{if .LicenseLink}}- [{{.LicenseName}}]({{.LicenseLink}}){{else}}- {{.LicenseName}}{{end}}{{end}}
{{end}}`
