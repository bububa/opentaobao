package mos

// CallDispatcherResponse 
/* model for simplify = false
type CallDispatcherResponse struct {

    // parentNo
    
    ParentNo   string `json:"parent_no,omitempty"`
    

    // respList
    
    RespList  struct {
        CallDispatcherRespDo  []CallDispatcherRespDo `json:"call_dispatcher_resp_do,omitempty"`
    } `json:"resp_list,omitempty"`
    

}
*/

// CallDispatcherResponse 
type CallDispatcherResponse struct {

    // parentNo
    ParentNo   string `json:"parent_no,omitempty"`

    // respList
    RespList   []CallDispatcherRespDo `json:"resp_list,omitempty"`

}
