package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkOrderDetailsGetAPIResponse 淘宝客-推广者-所有订单查询 API返回值
// taobao.tbk.order.details.get
//
// 淘宝客推广带来的所有拍下付款的正向订单明细报表。
type TaobaoTbkOrderDetailsGetAPIResponse struct {
	model.CommonResponse
	TaobaoTbkOrderDetailsGetAPIResponseModel
}

// TaobaoTbkOrderDetailsGetAPIResponseModel is 淘宝客-推广者-所有订单查询 成功返回结果
type TaobaoTbkOrderDetailsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_order_details_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// PublisherOrderDto
	Data *OrderPage `json:"data,omitempty" xml:"data,omitempty"`
}
