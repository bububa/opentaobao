package mos

// CallDispatcherResponse 
type CallDispatcherResponse struct {

    // parentNo
    ParentNo   string `json:"parent_no,omitempty"`

    // respList
    RespList   []CallDispatcherRespDo `json:"resp_list,omitempty"`

}
