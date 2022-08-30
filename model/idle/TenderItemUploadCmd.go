package idle

// TenderItemUploadCmd 结构体
type TenderItemUploadCmd struct {
	// 商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 质检报告
	Report string `json:"report,omitempty" xml:"report,omitempty"`
	// 上拍场次日期
	ScheduleDate string `json:"schedule_date,omitempty" xml:"schedule_date,omitempty"`
	// 外部码
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 场次编码（两位数值）
	ScheduleNumber string `json:"schedule_number,omitempty" xml:"schedule_number,omitempty"`
	// 动作类型:add：上传，update：更新
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 货源类型，RECYCLE（回收）, OTHER（其他，服务商自有）
	SourceType string `json:"source_type,omitempty" xml:"source_type,omitempty"`
}
