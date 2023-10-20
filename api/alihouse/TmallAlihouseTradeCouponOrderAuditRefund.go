package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// TmallAlihouseTradeCouponOrderAuditRefund ETC审核电商券退款
// tmall.alihouse.trade.coupon.order.audit.refund
//
// ETC审核电商券退款
func TmallAlihouseTradeCouponOrderAuditRefund(clt *core.SDKClient, req *alihouse.TmallAlihouseTradeCouponOrderAuditRefundAPIRequest, resp *alihouse.TmallAlihouseTradeCouponOrderAuditRefundAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
