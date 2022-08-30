package user

// ThirdPartyClientDataParams 结构体
type ThirdPartyClientDataParams struct {
	// 扩展字段,这里可支持扩展，但是需要报备
	Ext string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 数据生成时间
	GenerationTime string `json:"generation_time,omitempty" xml:"generation_time,omitempty"`
	// 电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 录入信息
	RecordContents string `json:"record_contents,omitempty" xml:"record_contents,omitempty"`
	// 外部数据记录唯一id
	RecordId string `json:"record_id,omitempty" xml:"record_id,omitempty"`
	// 录入人
	Recorder string `json:"recorder,omitempty" xml:"recorder,omitempty"`
	// 服务代码
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 服务类型：电话 PHONE, 预约上门APPOINTMENT_TO_DOOR, 未预约上门 NO_APPOINTMENT_TO_DOOR,  培训服务 TRAIN, 在线拜访 VISIT_ONLINE, 其他 OTHER;
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 客户数据
	Client *Client `json:"client,omitempty" xml:"client,omitempty"`
}
