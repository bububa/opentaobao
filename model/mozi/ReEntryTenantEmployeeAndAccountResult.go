package mozi

// ReEntryTenantEmployeeAndAccountResult 结构体
type ReEntryTenantEmployeeAndAccountResult struct {
	// 返回的人员和账号的绑定对象
	Datas []EmployeeAccount `json:"datas,omitempty" xml:"datas>employee_account,omitempty"`
	// 返回状态描述
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// 返回附加信息
	ResponseMetaData string `json:"response_meta_data,omitempty" xml:"response_meta_data,omitempty"`
	// 返回的状态码
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 请求的UUID
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
