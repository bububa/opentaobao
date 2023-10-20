package idleisv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvPvQueryAPIResponse 查询pv属性 API返回值
// alibaba.idle.isv.pv.query
//
// 查询pv属性
type AlibabaIdleIsvPvQueryAPIResponse struct {
	model.CommonResponse
	AlibabaIdleIsvPvQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleIsvPvQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleIsvPvQueryAPIResponseModel).Reset()
}

// AlibabaIdleIsvPvQueryAPIResponseModel is 查询pv属性 成功返回结果
type AlibabaIdleIsvPvQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_pv_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleIsvPvQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleIsvPvQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvPvQueryAPIResponse)
	},
}

// GetAlibabaIdleIsvPvQueryAPIResponse 从 sync.Pool 获取 AlibabaIdleIsvPvQueryAPIResponse
func GetAlibabaIdleIsvPvQueryAPIResponse() *AlibabaIdleIsvPvQueryAPIResponse {
	return poolAlibabaIdleIsvPvQueryAPIResponse.Get().(*AlibabaIdleIsvPvQueryAPIResponse)
}

// ReleaseAlibabaIdleIsvPvQueryAPIResponse 将 AlibabaIdleIsvPvQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleIsvPvQueryAPIResponse(v *AlibabaIdleIsvPvQueryAPIResponse) {
	v.Reset()
	poolAlibabaIdleIsvPvQueryAPIResponse.Put(v)
}
