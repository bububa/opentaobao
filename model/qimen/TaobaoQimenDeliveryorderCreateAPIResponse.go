package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenDeliveryorderCreateAPIResponse 发货单创建接口 API返回值
// taobao.qimen.deliveryorder.create
//
// taobao.qimen.deliveryorder.create
type TaobaoQimenDeliveryorderCreateAPIResponse struct {
	model.CommonResponse
	TaobaoQimenDeliveryorderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenDeliveryorderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenDeliveryorderCreateAPIResponseModel).Reset()
}

// TaobaoQimenDeliveryorderCreateAPIResponseModel is 发货单创建接口 成功返回结果
type TaobaoQimenDeliveryorderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_deliveryorder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应字段
	Response *DeliveryOrderCreateResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenDeliveryorderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenDeliveryorderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenDeliveryorderCreateAPIResponse)
	},
}

// GetTaobaoQimenDeliveryorderCreateAPIResponse 从 sync.Pool 获取 TaobaoQimenDeliveryorderCreateAPIResponse
func GetTaobaoQimenDeliveryorderCreateAPIResponse() *TaobaoQimenDeliveryorderCreateAPIResponse {
	return poolTaobaoQimenDeliveryorderCreateAPIResponse.Get().(*TaobaoQimenDeliveryorderCreateAPIResponse)
}

// ReleaseTaobaoQimenDeliveryorderCreateAPIResponse 将 TaobaoQimenDeliveryorderCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenDeliveryorderCreateAPIResponse(v *TaobaoQimenDeliveryorderCreateAPIResponse) {
	v.Reset()
	poolTaobaoQimenDeliveryorderCreateAPIResponse.Put(v)
}
