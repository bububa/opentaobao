package aesolution

// ProductSchemaDto 结构体
type ProductSchemaDto struct {
	// error code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// error message
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// schema
	Schema string `json:"schema,omitempty" xml:"schema,omitempty"`
	// success flag
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
