package tmallservice

// ServiceTaskPacket 
type ServiceTaskPacket struct {

    // 服务名字
    Name   string `json:"name,omitempty"`

    // 描述
    Desc   string `json:"desc,omitempty"`

    // 服务工单DO
    ServiceList   []ServiceTaskDO `json:"service_list,omitempty"`

}
