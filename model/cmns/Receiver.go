package cmns

// Receiver 
/* model for simplify = false
type Receiver struct {

    // deviceToken值，最多1000个
    
    Data  struct {
        String  []string `json:"string,omitempty"`
    } `json:"data,omitempty"`
    

    // 只能填写deviceToken
    
    Type   string `json:"type,omitempty"`
    

    // deviceToken值，最多1000个
    
    DataList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"data_list,omitempty"`
    

}
*/

// Receiver 
type Receiver struct {

    // deviceToken值，最多1000个
    Data   []string `json:"data,omitempty"`

    // 只能填写deviceToken
    Type   string `json:"type,omitempty"`

    // deviceToken值，最多1000个
    DataList   []string `json:"data_list,omitempty"`

}
