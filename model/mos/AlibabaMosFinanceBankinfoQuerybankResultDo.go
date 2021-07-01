package mos

// AlibabaMosFinanceBankinfoQuerybankResultDo 结构体
type AlibabaMosFinanceBankinfoQuerybankResultDo struct {
	// 总量
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 扩展
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 全链路追踪id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 描述
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Data *SupplierBankInfoDto `json:"data,omitempty" xml:"data,omitempty"`
	// 成功
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 返回值
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
