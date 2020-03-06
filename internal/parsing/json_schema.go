package parsing

// available Draft URIs in JSON Schema
type DraftVersions = string

// available DataTypes in JSON Schema
type DataTypes = string

// available String Formats in JSON Schema
type StringFormat = string

const (
	DraftV1 DraftVersions = "http://json-schema.org/draft-01/schema#"
	DraftV2 DraftVersions = "http://json-schema.org/draft-02/schema#"
	DraftV3 DraftVersions = "http://json-schema.org/draft-03/schema#"
	DraftV4 DraftVersions = "http://json-schema.org/draft-04/schema#"
	DraftV5 DraftVersions = "http://json-schema.org/draft-05/schema#"
	DraftV6 DraftVersions = "http://json-schema.org/draft-06/schema#"
	DraftV7 DraftVersions = "http://json-schema.org/draft-07/schema#"
	DraftV8 DraftVersions = "http://json-schema.org/draft-08/schema#"
	DraftV9 DraftVersions = "http://json-schema.org/draft-09/schema#"
)

const (
	Bool    DataTypes = "boolean"
	Integer           = "integer"
	Number            = "number"
	String            = "string"
	Object            = "object"
	Array             = "array"
	Null              = "null"
	Any               = "any"
)

const (
	Date              StringFormat = "date"
	Time                           = "time"
	DateTime                       = "date-time"
	RegularExpression              = "regex"
	Phone                          = "phone"
	Uri                            = "uri"
	Email                          = "email"
)
