package examination

import (
	"sync"
)

// IsvGetReportFileCodeResponse 结构体
type IsvGetReportFileCodeResponse struct {
	// 报告文件查看验证码
	IsvCode string `json:"isv_code,omitempty" xml:"isv_code,omitempty"`
}

var poolIsvGetReportFileCodeResponse = sync.Pool{
	New: func() any {
		return new(IsvGetReportFileCodeResponse)
	},
}

// GetIsvGetReportFileCodeResponse() 从对象池中获取IsvGetReportFileCodeResponse
func GetIsvGetReportFileCodeResponse() *IsvGetReportFileCodeResponse {
	return poolIsvGetReportFileCodeResponse.Get().(*IsvGetReportFileCodeResponse)
}

// ReleaseIsvGetReportFileCodeResponse 释放IsvGetReportFileCodeResponse
func ReleaseIsvGetReportFileCodeResponse(v *IsvGetReportFileCodeResponse) {
	v.IsvCode = ""
	poolIsvGetReportFileCodeResponse.Put(v)
}
