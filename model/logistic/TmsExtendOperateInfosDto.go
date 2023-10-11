package logistic

// TmsExtendOperateInfosDto 结构体
type TmsExtendOperateInfosDto struct {
	// 操作类型（枚举）： UpdateAddress-服务商修改地址
	OperateType string `json:"operate_type,omitempty" xml:"operate_type,omitempty"`
	// 操作时间(YYYY-MM-DD HH:MM:SS)
	OperateTime string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 内容
	OperateDetail string `json:"operate_detail,omitempty" xml:"operate_detail,omitempty"`
}
