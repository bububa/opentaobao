package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallalihousetradecouponrefundorderqueryAPIResponse 查询电商券履约退款单 API返回值
// tmall.alihouse.trade.coupon.refund.order.query
//
// 查询电商券履约退款单
type TmallalihousetradecouponrefundorderqueryAPIResponse struct {
	model.CommonResponse
	TmallalihousetradecouponrefundorderqueryAPIResponseModel
}

// TmallalihousetradecouponrefundorderqueryAPIResponseModel is 查询电商券履约退款单 成功返回结果
type TmallalihousetradecouponrefundorderqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_alihouse_trade_coupon_refund_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallalihousetradecouponrefundorderqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
