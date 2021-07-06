package security

// RpGrade 结构体
type RpGrade struct {
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// level
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
}
