package tmallservice

// ServiceTaskPacket 结构体
type ServiceTaskPacket struct {
	// 服务工单DO
	ServiceList []ServiceTaskDo `json:"service_list,omitempty" xml:"service_list>service_task_do,omitempty"`
	// 服务名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
}
