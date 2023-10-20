package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRecycleOrderPerformAPIResponse 回收订单履约V2 API返回值
// alibaba.idle.recycle.order.perform
//
// 闲鱼回收业务中,外部服务商作为买家 需要驱动交易节点变化
type AlibabaIdleRecycleOrderPerformAPIResponse struct {
	model.CommonResponse
	AlibabaIdleRecycleOrderPerformAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleRecycleOrderPerformAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleRecycleOrderPerformAPIResponseModel).Reset()
}

// AlibabaIdleRecycleOrderPerformAPIResponseModel is 回收订单履约V2 成功返回结果
type AlibabaIdleRecycleOrderPerformAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_recycle_order_perform_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleRecycleOrderPerformAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaIdleRecycleOrderPerformAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRecycleOrderPerformAPIResponse)
	},
}

// GetAlibabaIdleRecycleOrderPerformAPIResponse 从 sync.Pool 获取 AlibabaIdleRecycleOrderPerformAPIResponse
func GetAlibabaIdleRecycleOrderPerformAPIResponse() *AlibabaIdleRecycleOrderPerformAPIResponse {
	return poolAlibabaIdleRecycleOrderPerformAPIResponse.Get().(*AlibabaIdleRecycleOrderPerformAPIResponse)
}

// ReleaseAlibabaIdleRecycleOrderPerformAPIResponse 将 AlibabaIdleRecycleOrderPerformAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleRecycleOrderPerformAPIResponse(v *AlibabaIdleRecycleOrderPerformAPIResponse) {
	v.Reset()
	poolAlibabaIdleRecycleOrderPerformAPIResponse.Put(v)
}
