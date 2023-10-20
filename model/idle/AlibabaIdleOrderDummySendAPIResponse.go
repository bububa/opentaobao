package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleOrderDummySendAPIResponse 闲鱼无需物流发货 API返回值
// alibaba.idle.order.dummy.send
//
// 适用于电子卡券等虚拟商品不需要物流的商品发货。
type AlibabaIdleOrderDummySendAPIResponse struct {
	model.CommonResponse
	AlibabaIdleOrderDummySendAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleOrderDummySendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleOrderDummySendAPIResponseModel).Reset()
}

// AlibabaIdleOrderDummySendAPIResponseModel is 闲鱼无需物流发货 成功返回结果
type AlibabaIdleOrderDummySendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_order_dummy_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleOrderDummySendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleOrderDummySendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleOrderDummySendAPIResponse)
	},
}

// GetAlibabaIdleOrderDummySendAPIResponse 从 sync.Pool 获取 AlibabaIdleOrderDummySendAPIResponse
func GetAlibabaIdleOrderDummySendAPIResponse() *AlibabaIdleOrderDummySendAPIResponse {
	return poolAlibabaIdleOrderDummySendAPIResponse.Get().(*AlibabaIdleOrderDummySendAPIResponse)
}

// ReleaseAlibabaIdleOrderDummySendAPIResponse 将 AlibabaIdleOrderDummySendAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleOrderDummySendAPIResponse(v *AlibabaIdleOrderDummySendAPIResponse) {
	v.Reset()
	poolAlibabaIdleOrderDummySendAPIResponse.Put(v)
}
