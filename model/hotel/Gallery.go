package hotel

// Gallery 
type Gallery struct {

    // id
    Id   int64 `json:"id,omitempty"`

    // name
    Name   string `json:"name,omitempty"`

    // pictures
    Pictures   []String `json:"pictures,omitempty"`

}
