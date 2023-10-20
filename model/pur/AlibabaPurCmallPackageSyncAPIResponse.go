package pur

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurCmallPackageSyncAPIResponse 套餐同步 API返回值
// alibaba.pur.cmall.package.sync
//
// 套餐同步
type AlibabaPurCmallPackageSyncAPIResponse struct {
	model.CommonResponse
	AlibabaPurCmallPackageSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaPurCmallPackageSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaPurCmallPackageSyncAPIResponseModel).Reset()
}

// AlibabaPurCmallPackageSyncAPIResponseModel is 套餐同步 成功返回结果
type AlibabaPurCmallPackageSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pur_cmall_package_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaPurCmallPackageSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaPurCmallPackageSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaPurCmallPackageSyncAPIResponse)
	},
}

// GetAlibabaPurCmallPackageSyncAPIResponse 从 sync.Pool 获取 AlibabaPurCmallPackageSyncAPIResponse
func GetAlibabaPurCmallPackageSyncAPIResponse() *AlibabaPurCmallPackageSyncAPIResponse {
	return poolAlibabaPurCmallPackageSyncAPIResponse.Get().(*AlibabaPurCmallPackageSyncAPIResponse)
}

// ReleaseAlibabaPurCmallPackageSyncAPIResponse 将 AlibabaPurCmallPackageSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaPurCmallPackageSyncAPIResponse(v *AlibabaPurCmallPackageSyncAPIResponse) {
	v.Reset()
	poolAlibabaPurCmallPackageSyncAPIResponse.Put(v)
}
