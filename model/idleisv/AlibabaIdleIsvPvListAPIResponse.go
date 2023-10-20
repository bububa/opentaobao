package idleisv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvPvListAPIResponse 闲鱼已验货pv查询 API返回值
// alibaba.idle.isv.pv.list
//
// 根据闲鱼渠道类目查询对应的品牌和型号清单，供服务商进行选择
type AlibabaIdleIsvPvListAPIResponse struct {
	model.CommonResponse
	AlibabaIdleIsvPvListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleIsvPvListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleIsvPvListAPIResponseModel).Reset()
}

// AlibabaIdleIsvPvListAPIResponseModel is 闲鱼已验货pv查询 成功返回结果
type AlibabaIdleIsvPvListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_pv_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaIdleIsvPvListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleIsvPvListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleIsvPvListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvPvListAPIResponse)
	},
}

// GetAlibabaIdleIsvPvListAPIResponse 从 sync.Pool 获取 AlibabaIdleIsvPvListAPIResponse
func GetAlibabaIdleIsvPvListAPIResponse() *AlibabaIdleIsvPvListAPIResponse {
	return poolAlibabaIdleIsvPvListAPIResponse.Get().(*AlibabaIdleIsvPvListAPIResponse)
}

// ReleaseAlibabaIdleIsvPvListAPIResponse 将 AlibabaIdleIsvPvListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleIsvPvListAPIResponse(v *AlibabaIdleIsvPvListAPIResponse) {
	v.Reset()
	poolAlibabaIdleIsvPvListAPIResponse.Put(v)
}
