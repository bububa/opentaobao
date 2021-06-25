package cmns

// Receiver 
type Receiver struct {

    // deviceToken值，最多1000个
    Data   []String `json:"data,omitempty"`

    // 只能填写deviceToken
    Type   string `json:"type,omitempty"`

    // deviceToken值，最多1000个
    DataList   []String `json:"data_list,omitempty"`

}
