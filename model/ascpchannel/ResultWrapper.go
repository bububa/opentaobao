package ascpchannel

// ResultWrapper 结构体
type ResultWrapper struct {
	// 响应数据
	Datas []Datas `json:"datas,omitempty" xml:"datas>datas,omitempty"`
	// 系统自动生成
	Data []AlibabaascpaicsupplieraicinventorynegativesaleinvalidateData `json:"data,omitempty" xml:"data>alibabaascpaicsupplieraicinventorynegativesaleinvalidate_data,omitempty"`
	// 返回内容
	DataList []AlibabaascpchannelmainrefundcreateData `json:"data_list,omitempty" xml:"data_list>alibabaascpchannelmainrefundcreate_data,omitempty"`
	// 错误描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误码
	Error string `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 执行结果
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
	// 是否需要重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}
