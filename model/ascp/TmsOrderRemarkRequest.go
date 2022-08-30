package ascp

// TmsOrderRemarkRequest 结构体
type TmsOrderRemarkRequest struct {
	// 商家id
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 运单号
	ExpressCode string `json:"express_code,omitempty" xml:"express_code,omitempty"`
	// 业务请求ID
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 行业
	Industry string `json:"industry,omitempty" xml:"industry,omitempty"`
	// 服务商code
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 操作人
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 扩展字段
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 业务请求时间戳
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}
