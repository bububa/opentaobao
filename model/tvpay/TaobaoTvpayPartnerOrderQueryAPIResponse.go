package tvpay

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotvpaypartnerorderqueryAPIResponse 商户查询订单 API返回值
// taobao.tvpay.partner.order.query
//
// 给商户提供的查询订单状态的API
type TaobaotvpaypartnerorderqueryAPIResponse struct {
	model.CommonResponse
	TaobaotvpaypartnerorderqueryAPIResponseModel
}

// TaobaotvpaypartnerorderqueryAPIResponseModel is 商户查询订单 成功返回结果
type TaobaotvpaypartnerorderqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tvpay_partner_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Top返回对象
	Result *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
