package logistic

// TaobaoWlbImportThreeplResourceGetTopResult 结构体
type TaobaoWlbImportThreeplResourceGetTopResult struct {
	// 资源列表
	Resources []ThreePlConsignResourceDto `json:"resources,omitempty" xml:"resources>three_pl_consign_resource_dto,omitempty"`
	// 错误信息
	SubErrorMsg string `json:"sub_error_msg,omitempty" xml:"sub_error_msg,omitempty"`
	// 错误代码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误代码
	SubErrorCode string `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
