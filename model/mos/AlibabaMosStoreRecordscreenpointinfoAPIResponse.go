package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosStoreRecordscreenpointinfoAPIResponse 云屏埋点数据记录接口 API返回值
// alibaba.mos.store.recordscreenpointinfo
//
// 记录云屏埋点数据
type AlibabaMosStoreRecordscreenpointinfoAPIResponse struct {
	model.CommonResponse
	AlibabaMosStoreRecordscreenpointinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosStoreRecordscreenpointinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosStoreRecordscreenpointinfoAPIResponseModel).Reset()
}

// AlibabaMosStoreRecordscreenpointinfoAPIResponseModel is 云屏埋点数据记录接口 成功返回结果
type AlibabaMosStoreRecordscreenpointinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_store_recordscreenpointinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaMosStoreRecordscreenpointinfoResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosStoreRecordscreenpointinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMosStoreRecordscreenpointinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosStoreRecordscreenpointinfoAPIResponse)
	},
}

// GetAlibabaMosStoreRecordscreenpointinfoAPIResponse 从 sync.Pool 获取 AlibabaMosStoreRecordscreenpointinfoAPIResponse
func GetAlibabaMosStoreRecordscreenpointinfoAPIResponse() *AlibabaMosStoreRecordscreenpointinfoAPIResponse {
	return poolAlibabaMosStoreRecordscreenpointinfoAPIResponse.Get().(*AlibabaMosStoreRecordscreenpointinfoAPIResponse)
}

// ReleaseAlibabaMosStoreRecordscreenpointinfoAPIResponse 将 AlibabaMosStoreRecordscreenpointinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosStoreRecordscreenpointinfoAPIResponse(v *AlibabaMosStoreRecordscreenpointinfoAPIResponse) {
	v.Reset()
	poolAlibabaMosStoreRecordscreenpointinfoAPIResponse.Put(v)
}
