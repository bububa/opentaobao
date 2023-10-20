package alilabs

// AlibabaailabuserprofilegetResult 结构体
type AlibabaailabuserprofilegetResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 数据model
	Result *BasicUserInfo `json:"result,omitempty" xml:"result,omitempty"`
	// 200 成功，其他 失败
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
}
