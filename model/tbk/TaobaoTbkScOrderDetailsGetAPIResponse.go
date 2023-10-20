package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScOrderDetailsGetAPIResponse 淘宝客-服务商-所有订单查询 API返回值
// taobao.tbk.sc.order.details.get
//
// 服务商使用。可通过该接口查询推广者账号下对应的推广订单明细。
type TaobaoTbkScOrderDetailsGetAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScOrderDetailsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkScOrderDetailsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkScOrderDetailsGetAPIResponseModel).Reset()
}

// TaobaoTbkScOrderDetailsGetAPIResponseModel is 淘宝客-服务商-所有订单查询 成功返回结果
type TaobaoTbkScOrderDetailsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_order_details_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// PublisherOrderDto
	Data *OrderPage `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkScOrderDetailsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoTbkScOrderDetailsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScOrderDetailsGetAPIResponse)
	},
}

// GetTaobaoTbkScOrderDetailsGetAPIResponse 从 sync.Pool 获取 TaobaoTbkScOrderDetailsGetAPIResponse
func GetTaobaoTbkScOrderDetailsGetAPIResponse() *TaobaoTbkScOrderDetailsGetAPIResponse {
	return poolTaobaoTbkScOrderDetailsGetAPIResponse.Get().(*TaobaoTbkScOrderDetailsGetAPIResponse)
}

// ReleaseTaobaoTbkScOrderDetailsGetAPIResponse 将 TaobaoTbkScOrderDetailsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkScOrderDetailsGetAPIResponse(v *TaobaoTbkScOrderDetailsGetAPIResponse) {
	v.Reset()
	poolTaobaoTbkScOrderDetailsGetAPIResponse.Put(v)
}
