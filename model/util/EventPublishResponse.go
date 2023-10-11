package util

// EventPublishResponse 结构体
type EventPublishResponse struct {
	// 流程返回值列表
	ResponseList []ProcessResponse `json:"response_list,omitempty" xml:"response_list>process_response,omitempty"`
	// traceId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 是否触发工作流
	ExecuteProcess bool `json:"execute_process,omitempty" xml:"execute_process,omitempty"`
}
