package lstpos

// AlibabalstposopeninventorysyncinventorydataResultDto 结构体
type AlibabalstposopeninventorysyncinventorydataResultDto struct {
	// 接口具体返回的业务数据对象
	ModuleList []ErrorResult `json:"module_list,omitempty" xml:"module_list>error_result,omitempty"`
	// 错误消息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 接口调用是否成功 true:调用成功 false:调用失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
