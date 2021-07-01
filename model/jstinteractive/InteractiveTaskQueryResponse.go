package jstinteractive

// InteractiveTaskQueryResponse 结构体
type InteractiveTaskQueryResponse struct {
	// 任务列表
	TaskList []InteractiveTask `json:"task_list,omitempty" xml:"task_list>interactive_task,omitempty"`
}
