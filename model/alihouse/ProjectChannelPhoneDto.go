package alihouse

// ProjectChannelPhoneDto 结构体
type ProjectChannelPhoneDto struct {
	// 楼盘outerid
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 渠道电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 批次号
	BatchNo string `json:"batch_no,omitempty" xml:"batch_no,omitempty"`
	// 渠道类型 1-高德
	Channel int64 `json:"channel,omitempty" xml:"channel,omitempty"`
	// 是否删除 1-已删除 0-未删除
	IsDeleted int64 `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
}
