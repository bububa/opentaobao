package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Tmallalihousetradecouponorderauditrefund ETC审核电商券退款
// tmall.alihouse.trade.coupon.order.audit.refund
//
// ETC审核电商券退款
func Tmallalihousetradecouponorderauditrefund(clt *core.SDKClient, req *alihouse.TmallalihousetradecouponorderauditrefundAPIRequest, session string) (*alihouse.TmallalihousetradecouponorderauditrefundAPIResponse, error) {
	var resp alihouse.TmallalihousetradecouponorderauditrefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
