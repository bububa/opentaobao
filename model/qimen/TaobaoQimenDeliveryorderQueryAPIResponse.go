package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenDeliveryorderQueryAPIResponse 发货单查询接口 API返回值
// taobao.qimen.deliveryorder.query
//
// ERP调用奇门的发货单查询接口，查询发货单详情
type TaobaoQimenDeliveryorderQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenDeliveryorderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenDeliveryorderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenDeliveryorderQueryAPIResponseModel).Reset()
}

// TaobaoQimenDeliveryorderQueryAPIResponseModel is 发货单查询接口 成功返回结果
type TaobaoQimenDeliveryorderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_deliveryorder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *DeliveryOrderQueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenDeliveryorderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenDeliveryorderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenDeliveryorderQueryAPIResponse)
	},
}

// GetTaobaoQimenDeliveryorderQueryAPIResponse 从 sync.Pool 获取 TaobaoQimenDeliveryorderQueryAPIResponse
func GetTaobaoQimenDeliveryorderQueryAPIResponse() *TaobaoQimenDeliveryorderQueryAPIResponse {
	return poolTaobaoQimenDeliveryorderQueryAPIResponse.Get().(*TaobaoQimenDeliveryorderQueryAPIResponse)
}

// ReleaseTaobaoQimenDeliveryorderQueryAPIResponse 将 TaobaoQimenDeliveryorderQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenDeliveryorderQueryAPIResponse(v *TaobaoQimenDeliveryorderQueryAPIResponse) {
	v.Reset()
	poolTaobaoQimenDeliveryorderQueryAPIResponse.Put(v)
}
