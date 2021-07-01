package tvupadmin

// AdvertScheduleDo 结构体
type AdvertScheduleDo struct {
	// 查询类型
	Range int64 `json:"range,omitempty" xml:"range,omitempty"`
	// 播控ID
	BcpId int64 `json:"bcp_id,omitempty" xml:"bcp_id,omitempty"`
	// 设备型号
	DeviceModel string `json:"device_model,omitempty" xml:"device_model,omitempty"`
	// UUID
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 广告类型
	SiteType int64 `json:"site_type,omitempty" xml:"site_type,omitempty"`
	// 开始时间
	Start string `json:"start,omitempty" xml:"start,omitempty"`
	// 结束时间
	End string `json:"end,omitempty" xml:"end,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
