package attribute

type attributionType string

type attribution struct {
	template string
}

// licenses binds all the licenses to generate together.
var attributions = map[attributionType]string{
	attributionTypeShort: attributionShort,
}
