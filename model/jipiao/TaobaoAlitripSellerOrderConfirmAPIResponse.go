package jipiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripSellerOrderConfirmAPIResponse 代理商确认机票订单接口 API返回值
// taobao.alitrip.seller.order.confirm
//
// 此接口用于代理商确认机票订单。
type TaobaoAlitripSellerOrderConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripSellerOrderConfirmAPIResponseModel
}

// TaobaoAlitripSellerOrderConfirmAPIResponseModel is 代理商确认机票订单接口 成功返回结果
type TaobaoAlitripSellerOrderConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_seller_order_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *TripOrderConfirmResultVo `json:"result,omitempty" xml:"result,omitempty"`
}
