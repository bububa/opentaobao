package tmallservice

// DivisionDto 
/* model for simplify = false
type DivisionDto struct {

    // 1
    
    DivisionNames  struct {
        String  []string `json:"string,omitempty"`
    } `json:"division_names,omitempty"`
    

}
*/

// DivisionDto 
type DivisionDto struct {

    // 1
    DivisionNames   []string `json:"division_names,omitempty"`

}
