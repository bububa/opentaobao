package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// TmallAlihouseTradeCouponRefundOrderQuery 查询电商券履约退款单
// tmall.alihouse.trade.coupon.refund.order.query
//
// 查询电商券履约退款单
func TmallAlihouseTradeCouponRefundOrderQuery(clt *core.SDKClient, req *alihouse.TmallAlihouseTradeCouponRefundOrderQueryAPIRequest, resp *alihouse.TmallAlihouseTradeCouponRefundOrderQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
