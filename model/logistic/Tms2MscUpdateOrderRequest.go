package logistic

// Tms2MscUpdateOrderRequest 结构体
type Tms2MscUpdateOrderRequest struct {
	// 电联信息（组）
	PhoneCallInfos []TmsPhoneCallInfoDto `json:"phone_call_infos,omitempty" xml:"phone_call_infos>tms_phone_call_info_dto,omitempty"`
	// 业务类型
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 淘天物流服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 修改后的上门日期（YYYY-MM-DD）
	ServiceDate string `json:"service_date,omitempty" xml:"service_date,omitempty"`
	// 请求类型
	RequestType string `json:"request_type,omitempty" xml:"request_type,omitempty"`
	// 唯一标识单号
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 修改上门时间原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 修改后的上门时间段（HH:MM-HH:MM）
	ServiceTimeRange string `json:"service_time_range,omitempty" xml:"service_time_range,omitempty"`
	// 配资源编码
	DeliverCode string `json:"deliver_code,omitempty" xml:"deliver_code,omitempty"`
}
