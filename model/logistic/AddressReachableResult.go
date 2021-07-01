package logistic

// AddressReachableResult 结构体
type AddressReachableResult struct {
	// 物流公司编码ID
	PartnerId int64 `json:"partner_id,omitempty" xml:"partner_id,omitempty"`
	// 服务是否可达，    0 - 不可达    1 - 可达    2 - 不确定   3 - 未配置
	Reachable int64 `json:"reachable,omitempty" xml:"reachable,omitempty"`
	// 服务对应的数字编码，如揽派范围对应88
	ServiceType int64 `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 区域编码
	DivisionId int64 `json:"division_id,omitempty" xml:"division_id,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 物流公司代号
	PartnerCode string `json:"partner_code,omitempty" xml:"partner_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
