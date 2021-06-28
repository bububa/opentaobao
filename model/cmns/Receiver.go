package cmns

// Receiver 
type Receiver struct {

    // deviceToken值，最多1000个
    
    Data   []string `json:"data,omitempty" xml:"data>string,omitempty"`
    

    // 只能填写deviceToken
    
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    

    // deviceToken值，最多1000个
    
    DataList   []string `json:"data_list,omitempty" xml:"data_list>string,omitempty"`
    

}
