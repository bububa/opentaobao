package hotel

// Gallery 结构体
type Gallery struct {
	// pictures
	Pictures []string `json:"pictures,omitempty" xml:"pictures>string,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
