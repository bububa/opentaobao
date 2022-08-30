package wlbimports

// AppointmentOrderStatusResponse 结构体
type AppointmentOrderStatusResponse struct {
	// 预约单code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 状态code
	StatusCode string `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 状态名称
	StatusName string `json:"status_name,omitempty" xml:"status_name,omitempty"`
	// 更新时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 主键id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 商家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}
