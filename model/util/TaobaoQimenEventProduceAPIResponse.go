package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenEventProduceAPIResponse 发出奇门事件 API返回值
// taobao.qimen.event.produce
//
// 当订单被处理时，用于通知奇门系统。
type TaobaoQimenEventProduceAPIResponse struct {
	model.CommonResponse
	TaobaoQimenEventProduceAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenEventProduceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenEventProduceAPIResponseModel).Reset()
}

// TaobaoQimenEventProduceAPIResponseModel is 发出奇门事件 成功返回结果
type TaobaoQimenEventProduceAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_event_produce_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenEventProduceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoQimenEventProduceAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenEventProduceAPIResponse)
	},
}

// GetTaobaoQimenEventProduceAPIResponse 从 sync.Pool 获取 TaobaoQimenEventProduceAPIResponse
func GetTaobaoQimenEventProduceAPIResponse() *TaobaoQimenEventProduceAPIResponse {
	return poolTaobaoQimenEventProduceAPIResponse.Get().(*TaobaoQimenEventProduceAPIResponse)
}

// ReleaseTaobaoQimenEventProduceAPIResponse 将 TaobaoQimenEventProduceAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenEventProduceAPIResponse(v *TaobaoQimenEventProduceAPIResponse) {
	v.Reset()
	poolTaobaoQimenEventProduceAPIResponse.Put(v)
}
