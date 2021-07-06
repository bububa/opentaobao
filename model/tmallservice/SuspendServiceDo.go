package tmallservice

// SuspendServiceDo 结构体
type SuspendServiceDo struct {
	// api调用者
	Updater string `json:"updater,omitempty" xml:"updater,omitempty"`
	// 扩展字段
	Extension string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 无需安装的原因
	Comments string `json:"comments,omitempty" xml:"comments,omitempty"`
	// 工单类型，固定值1
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 买家user_id
	BuyerId int64 `json:"buyer_id,omitempty" xml:"buyer_id,omitempty"`
	// 更新时间得long值
	UpdateDate int64 `json:"update_date,omitempty" xml:"update_date,omitempty"`
	// 工单id
	WorkcardId int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
}
