package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleAppraiseOrderQueryAPIResponse 闲鱼验货宝订单详情查询 API返回值
// alibaba.idle.appraise.order.query
//
// 鉴定商调用该接口获取订单状态
type AlibabaIdleAppraiseOrderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaIdleAppraiseOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleAppraiseOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleAppraiseOrderQueryAPIResponseModel).Reset()
}

// AlibabaIdleAppraiseOrderQueryAPIResponseModel is 闲鱼验货宝订单详情查询 成功返回结果
type AlibabaIdleAppraiseOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_appraise_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaIdleAppraiseOrderQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleAppraiseOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleAppraiseOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleAppraiseOrderQueryAPIResponse)
	},
}

// GetAlibabaIdleAppraiseOrderQueryAPIResponse 从 sync.Pool 获取 AlibabaIdleAppraiseOrderQueryAPIResponse
func GetAlibabaIdleAppraiseOrderQueryAPIResponse() *AlibabaIdleAppraiseOrderQueryAPIResponse {
	return poolAlibabaIdleAppraiseOrderQueryAPIResponse.Get().(*AlibabaIdleAppraiseOrderQueryAPIResponse)
}

// ReleaseAlibabaIdleAppraiseOrderQueryAPIResponse 将 AlibabaIdleAppraiseOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleAppraiseOrderQueryAPIResponse(v *AlibabaIdleAppraiseOrderQueryAPIResponse) {
	v.Reset()
	poolAlibabaIdleAppraiseOrderQueryAPIResponse.Put(v)
}
