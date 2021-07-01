package tvpay

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTvpayOrderQueryAPIResponse
tv支付查询订单状态 API返回值
taobao.tvpay.order.query

tv支付查询订单状态 */
type TaobaoTvpayOrderQueryAPIResponse struct {
	model.CommonResponse
	TaobaoTvpayOrderQueryAPIResponseModel
}

// TaobaoTvpayOrderQueryAPIResponseModel is tv支付查询订单状态 成功返回结果
type TaobaoTvpayOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tvpay_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Top返回对象
	Result *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
