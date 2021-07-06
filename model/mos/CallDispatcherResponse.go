package mos

// CallDispatcherResponse 结构体
type CallDispatcherResponse struct {
	// respList
	RespList []CallDispatcherRespDo `json:"resp_list,omitempty" xml:"resp_list>call_dispatcher_resp_do,omitempty"`
	// parentNo
	ParentNo string `json:"parent_no,omitempty" xml:"parent_no,omitempty"`
}
