package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Tmallalihousetradecouponrefundorderquery 查询电商券履约退款单
// tmall.alihouse.trade.coupon.refund.order.query
//
// 查询电商券履约退款单
func Tmallalihousetradecouponrefundorderquery(clt *core.SDKClient, req *alihouse.TmallalihousetradecouponrefundorderqueryAPIRequest, session string) (*alihouse.TmallalihousetradecouponrefundorderqueryAPIResponse, error) {
	var resp alihouse.TmallalihousetradecouponrefundorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
