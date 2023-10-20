package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleAppraiseOrderPerformAPIResponse 闲鱼验货宝服务商订单履约 API返回值
// alibaba.idle.appraise.order.perform
//
// 闲鱼验货担保业务中,外部服务商作为鉴定方 需要驱动交易节点变化
type AlibabaIdleAppraiseOrderPerformAPIResponse struct {
	model.CommonResponse
	AlibabaIdleAppraiseOrderPerformAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleAppraiseOrderPerformAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleAppraiseOrderPerformAPIResponseModel).Reset()
}

// AlibabaIdleAppraiseOrderPerformAPIResponseModel is 闲鱼验货宝服务商订单履约 成功返回结果
type AlibabaIdleAppraiseOrderPerformAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_appraise_order_perform_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaIdleAppraiseOrderPerformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleAppraiseOrderPerformAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleAppraiseOrderPerformAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleAppraiseOrderPerformAPIResponse)
	},
}

// GetAlibabaIdleAppraiseOrderPerformAPIResponse 从 sync.Pool 获取 AlibabaIdleAppraiseOrderPerformAPIResponse
func GetAlibabaIdleAppraiseOrderPerformAPIResponse() *AlibabaIdleAppraiseOrderPerformAPIResponse {
	return poolAlibabaIdleAppraiseOrderPerformAPIResponse.Get().(*AlibabaIdleAppraiseOrderPerformAPIResponse)
}

// ReleaseAlibabaIdleAppraiseOrderPerformAPIResponse 将 AlibabaIdleAppraiseOrderPerformAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleAppraiseOrderPerformAPIResponse(v *AlibabaIdleAppraiseOrderPerformAPIResponse) {
	v.Reset()
	poolAlibabaIdleAppraiseOrderPerformAPIResponse.Put(v)
}
