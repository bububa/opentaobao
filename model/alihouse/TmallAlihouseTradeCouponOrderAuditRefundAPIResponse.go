package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallalihousetradecouponorderauditrefundAPIResponse ETC审核电商券退款 API返回值
// tmall.alihouse.trade.coupon.order.audit.refund
//
// ETC审核电商券退款
type TmallalihousetradecouponorderauditrefundAPIResponse struct {
	model.CommonResponse
	TmallalihousetradecouponorderauditrefundAPIResponseModel
}

// TmallalihousetradecouponorderauditrefundAPIResponseModel is ETC审核电商券退款 成功返回结果
type TmallalihousetradecouponorderauditrefundAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_alihouse_trade_coupon_order_audit_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallalihousetradecouponorderauditrefundResult `json:"result,omitempty" xml:"result,omitempty"`
}
