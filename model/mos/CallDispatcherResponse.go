package mos

// CallDispatcherResponse 
type CallDispatcherResponse struct {
    // parentNo
    ParentNo   string `json:"parent_no,omitempty" xml:"parent_no,omitempty"`
    // respList
    RespList   []CallDispatcherRespDO `json:"resp_list,omitempty" xml:"resp_list>call_dispatcher_resp_do,omitempty"`
}
