package logistic

import (
	"sync"
)

// CloudPrinterBindRequest 结构体
type CloudPrinterBindRequest struct {
	// 打印机 mac 地址
	Uid string `json:"uid,omitempty" xml:"uid,omitempty"`
	//  验证码
	VerifyCode string `json:"verify_code,omitempty" xml:"verify_code,omitempty"`
}

var poolCloudPrinterBindRequest = sync.Pool{
	New: func() any {
		return new(CloudPrinterBindRequest)
	},
}

// GetCloudPrinterBindRequest() 从对象池中获取CloudPrinterBindRequest
func GetCloudPrinterBindRequest() *CloudPrinterBindRequest {
	return poolCloudPrinterBindRequest.Get().(*CloudPrinterBindRequest)
}

// ReleaseCloudPrinterBindRequest 释放CloudPrinterBindRequest
func ReleaseCloudPrinterBindRequest(v *CloudPrinterBindRequest) {
	v.Uid = ""
	v.VerifyCode = ""
	poolCloudPrinterBindRequest.Put(v)
}
