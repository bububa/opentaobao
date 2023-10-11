package logistic

// TmsCourierRequest 结构体
type TmsCourierRequest struct {
	// 服务类型
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 单号
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 服务标识
	ServiceFlag string `json:"service_flag,omitempty" xml:"service_flag,omitempty"`
	// 快递公司资源编码
	TmsCpCode string `json:"tms_cp_code,omitempty" xml:"tms_cp_code,omitempty"`
	// 服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 配资源编码
	DeliveryCode string `json:"delivery_code,omitempty" xml:"delivery_code,omitempty"`
	// 更新类型 1-首次分配小件员 2-改派小件员
	UpdateType string `json:"update_type,omitempty" xml:"update_type,omitempty"`
	// 改派小件员原因 1-客服改派（上报超区、人工分单） 2-小件员转单 3-其他(update_type=2时，必填)
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 小件员信息
	TmsCourierInfo *TmsCourierInfoRequest `json:"tms_courier_info,omitempty" xml:"tms_courier_info,omitempty"`
}
