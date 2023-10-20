package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenOrderCancelAPIResponse 单据取消接口 API返回值
// taobao.qimen.order.cancel
//
// taobao.qimen.order.cancel
type TaobaoQimenOrderCancelAPIResponse struct {
	model.CommonResponse
	TaobaoQimenOrderCancelAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenOrderCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenOrderCancelAPIResponseModel).Reset()
}

// TaobaoQimenOrderCancelAPIResponseModel is 单据取消接口 成功返回结果
type TaobaoQimenOrderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *OrderCancelResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenOrderCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenOrderCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenOrderCancelAPIResponse)
	},
}

// GetTaobaoQimenOrderCancelAPIResponse 从 sync.Pool 获取 TaobaoQimenOrderCancelAPIResponse
func GetTaobaoQimenOrderCancelAPIResponse() *TaobaoQimenOrderCancelAPIResponse {
	return poolTaobaoQimenOrderCancelAPIResponse.Get().(*TaobaoQimenOrderCancelAPIResponse)
}

// ReleaseTaobaoQimenOrderCancelAPIResponse 将 TaobaoQimenOrderCancelAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenOrderCancelAPIResponse(v *TaobaoQimenOrderCancelAPIResponse) {
	v.Reset()
	poolTaobaoQimenOrderCancelAPIResponse.Put(v)
}
