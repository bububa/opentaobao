package hotel

// Gallery 
type Gallery struct {
    // id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // name
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // pictures
    Pictures   []string `json:"pictures,omitempty" xml:"pictures>string,omitempty"`
}
