package xml

import "strings"

type Tag struct {
	Name      string
	Omitempty bool
	Skip      bool
}

func ParseTag(tagString string) *Tag {
	if tagString == "-" {
		return &Tag{Skip: true}
	}
	res := &Tag{}
	tags := strings.Split(tagString, ",")
	if len(tags) == 0 {
		return res
	}
	res.Name = tags[0]
	for _, option := range tags[1:] {
		// WARNING: this TrimSpace causes ParseTag to accept
		// tags with whitespace in them. This differs from
		// encoding/xml's Marshal function, which panics
		// when an option contains whitespace (as of go1.27.1)
		optionTrimmed := strings.TrimSpace(option)
		if optionTrimmed == "omitempty" {
			res.Omitempty = true
		}
	}
	return res
}
