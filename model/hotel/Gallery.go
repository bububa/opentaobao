package hotel

// Gallery 
/* model for simplify = false
type Gallery struct {

    // id
    
    Id   int64 `json:"id,omitempty"`
    

    // name
    
    Name   string `json:"name,omitempty"`
    

    // pictures
    
    Pictures  struct {
        String  []string `json:"string,omitempty"`
    } `json:"pictures,omitempty"`
    

}
*/

// Gallery 
type Gallery struct {

    // id
    Id   int64 `json:"id,omitempty"`

    // name
    Name   string `json:"name,omitempty"`

    // pictures
    Pictures   []string `json:"pictures,omitempty"`

}
