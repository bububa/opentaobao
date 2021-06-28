package mos

// CallDispatcherResponse 
type CallDispatcherResponse struct {

    // parentNo
    
    ParentNo   string `json:"parent_no,omitempty" xml:"parent_no,omitempty"`
    

    // respList
    
    RespList   []CallDispatcherRespDo `json:"resp_list,omitempty" xml:"resp_list,omitempty"`
    

}
