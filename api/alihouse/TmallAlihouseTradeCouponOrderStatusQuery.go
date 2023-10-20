package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// TmallAlihouseTradeCouponOrderStatusQuery 查询电商券履约单状态
// tmall.alihouse.trade.coupon.order.status.query
//
// 查询电商券履约单状态
func TmallAlihouseTradeCouponOrderStatusQuery(clt *core.SDKClient, req *alihouse.TmallAlihouseTradeCouponOrderStatusQueryAPIRequest, session string) (*alihouse.TmallAlihouseTradeCouponOrderStatusQueryAPIResponse, error) {
	var resp alihouse.TmallAlihouseTradeCouponOrderStatusQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
