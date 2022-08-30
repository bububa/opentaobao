package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// TmallAlihouseTradeCouponRefundOrderQuery 查询电商券履约退款单
// tmall.alihouse.trade.coupon.refund.order.query
//
// 查询电商券履约退款单
func TmallAlihouseTradeCouponRefundOrderQuery(clt *core.SDKClient, req *alihouse.TmallAlihouseTradeCouponRefundOrderQueryAPIRequest, session string) (*alihouse.TmallAlihouseTradeCouponRefundOrderQueryAPIResponse, error) {
	var resp alihouse.TmallAlihouseTradeCouponRefundOrderQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
