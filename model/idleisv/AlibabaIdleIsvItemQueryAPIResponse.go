package idleisv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvItemQueryAPIResponse 服务商闲鱼商品查询 API返回值
// alibaba.idle.isv.item.query
//
// 服务商ISV闲鱼商品查询
type AlibabaIdleIsvItemQueryAPIResponse struct {
	model.CommonResponse
	AlibabaIdleIsvItemQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleIsvItemQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleIsvItemQueryAPIResponseModel).Reset()
}

// AlibabaIdleIsvItemQueryAPIResponseModel is 服务商闲鱼商品查询 成功返回结果
type AlibabaIdleIsvItemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_item_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果result
	Result *AlibabaIdleIsvItemQueryTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleIsvItemQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleIsvItemQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvItemQueryAPIResponse)
	},
}

// GetAlibabaIdleIsvItemQueryAPIResponse 从 sync.Pool 获取 AlibabaIdleIsvItemQueryAPIResponse
func GetAlibabaIdleIsvItemQueryAPIResponse() *AlibabaIdleIsvItemQueryAPIResponse {
	return poolAlibabaIdleIsvItemQueryAPIResponse.Get().(*AlibabaIdleIsvItemQueryAPIResponse)
}

// ReleaseAlibabaIdleIsvItemQueryAPIResponse 将 AlibabaIdleIsvItemQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleIsvItemQueryAPIResponse(v *AlibabaIdleIsvItemQueryAPIResponse) {
	v.Reset()
	poolAlibabaIdleIsvItemQueryAPIResponse.Put(v)
}
