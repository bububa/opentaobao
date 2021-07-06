package tvupadmin

// DicControlTaskDo 结构体
type DicControlTaskDo struct {
	// 操作者
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 任务描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 任务名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 设备型号
	Devices string `json:"devices,omitempty" xml:"devices,omitempty"`
	// 任务类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// apk id
	ApkId int64 `json:"apk_id,omitempty" xml:"apk_id,omitempty"`
	// 牌照方
	License int64 `json:"license,omitempty" xml:"license,omitempty"`
}
