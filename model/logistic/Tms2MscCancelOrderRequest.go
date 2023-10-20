package logistic

// Tms2mscCancelOrderRequest 结构体
type Tms2mscCancelOrderRequest struct {
	// 电联信息（组）
	PhoneCallInfos []TmsPhoneCallInfoDto `json:"phone_call_infos,omitempty" xml:"phone_call_infos>tms_phone_call_info_dto,omitempty"`
	// 服务类型
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 淘天物流服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 唯一标识单号
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 取消上门订单原因，枚举
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 配资源编码
	DeliverCode string `json:"deliver_code,omitempty" xml:"deliver_code,omitempty"`
}
