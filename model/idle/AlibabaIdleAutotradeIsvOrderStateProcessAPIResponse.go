package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleAutotradeIsvOrderStateProcessAPIResponse 闲鱼订单状态推进 API返回值
// alibaba.idle.autotrade.isv.order.state.process
//
// 闲鱼订单状态推进
type AlibabaIdleAutotradeIsvOrderStateProcessAPIResponse struct {
	model.CommonResponse
	AlibabaIdleAutotradeIsvOrderStateProcessAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleAutotradeIsvOrderStateProcessAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleAutotradeIsvOrderStateProcessAPIResponseModel).Reset()
}

// AlibabaIdleAutotradeIsvOrderStateProcessAPIResponseModel is 闲鱼订单状态推进 成功返回结果
type AlibabaIdleAutotradeIsvOrderStateProcessAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_autotrade_isv_order_state_process_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleAutotradeIsvOrderStateProcessAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleAutotradeIsvOrderStateProcessAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleAutotradeIsvOrderStateProcessAPIResponse)
	},
}

// GetAlibabaIdleAutotradeIsvOrderStateProcessAPIResponse 从 sync.Pool 获取 AlibabaIdleAutotradeIsvOrderStateProcessAPIResponse
func GetAlibabaIdleAutotradeIsvOrderStateProcessAPIResponse() *AlibabaIdleAutotradeIsvOrderStateProcessAPIResponse {
	return poolAlibabaIdleAutotradeIsvOrderStateProcessAPIResponse.Get().(*AlibabaIdleAutotradeIsvOrderStateProcessAPIResponse)
}

// ReleaseAlibabaIdleAutotradeIsvOrderStateProcessAPIResponse 将 AlibabaIdleAutotradeIsvOrderStateProcessAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleAutotradeIsvOrderStateProcessAPIResponse(v *AlibabaIdleAutotradeIsvOrderStateProcessAPIResponse) {
	v.Reset()
	poolAlibabaIdleAutotradeIsvOrderStateProcessAPIResponse.Put(v)
}
