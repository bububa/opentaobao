package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// TmallAlihouseTradeCouponOrderAuditRefund ETC审核电商券退款
// tmall.alihouse.trade.coupon.order.audit.refund
//
// ETC审核电商券退款
func TmallAlihouseTradeCouponOrderAuditRefund(clt *core.SDKClient, req *alihouse.TmallAlihouseTradeCouponOrderAuditRefundAPIRequest, session string) (*alihouse.TmallAlihouseTradeCouponOrderAuditRefundAPIResponse, error) {
	var resp alihouse.TmallAlihouseTradeCouponOrderAuditRefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
