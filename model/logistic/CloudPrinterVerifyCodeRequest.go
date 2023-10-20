package logistic

import (
	"sync"
)

// CloudPrinterVerifyCodeRequest 结构体
type CloudPrinterVerifyCodeRequest struct {
	//  打印机 id
	Uid string `json:"uid,omitempty" xml:"uid,omitempty"`
}

var poolCloudPrinterVerifyCodeRequest = sync.Pool{
	New: func() any {
		return new(CloudPrinterVerifyCodeRequest)
	},
}

// GetCloudPrinterVerifyCodeRequest() 从对象池中获取CloudPrinterVerifyCodeRequest
func GetCloudPrinterVerifyCodeRequest() *CloudPrinterVerifyCodeRequest {
	return poolCloudPrinterVerifyCodeRequest.Get().(*CloudPrinterVerifyCodeRequest)
}

// ReleaseCloudPrinterVerifyCodeRequest 释放CloudPrinterVerifyCodeRequest
func ReleaseCloudPrinterVerifyCodeRequest(v *CloudPrinterVerifyCodeRequest) {
	v.Uid = ""
	poolCloudPrinterVerifyCodeRequest.Put(v)
}
