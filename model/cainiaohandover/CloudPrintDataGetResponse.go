package cainiaohandover

import (
	"sync"
)

// CloudPrintDataGetResponse 结构体
type CloudPrintDataGetResponse struct {
	// 面单云打印数据
	PrintData string `json:"print_data,omitempty" xml:"print_data,omitempty"`
	// 面单云打印数据MD5加密串
	PrintDataMd5 string `json:"print_data_md5,omitempty" xml:"print_data_md5,omitempty"`
}

var poolCloudPrintDataGetResponse = sync.Pool{
	New: func() any {
		return new(CloudPrintDataGetResponse)
	},
}

// GetCloudPrintDataGetResponse() 从对象池中获取CloudPrintDataGetResponse
func GetCloudPrintDataGetResponse() *CloudPrintDataGetResponse {
	return poolCloudPrintDataGetResponse.Get().(*CloudPrintDataGetResponse)
}

// ReleaseCloudPrintDataGetResponse 释放CloudPrintDataGetResponse
func ReleaseCloudPrintDataGetResponse(v *CloudPrintDataGetResponse) {
	v.PrintData = ""
	v.PrintDataMd5 = ""
	poolCloudPrintDataGetResponse.Put(v)
}
