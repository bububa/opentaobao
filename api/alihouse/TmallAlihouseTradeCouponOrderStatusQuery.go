package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// TmallAlihouseTradeCouponOrderStatusQuery 查询电商券履约单状态
// tmall.alihouse.trade.coupon.order.status.query
//
// 查询电商券履约单状态
func TmallAlihouseTradeCouponOrderStatusQuery(clt *core.SDKClient, req *alihouse.TmallAlihouseTradeCouponOrderStatusQueryAPIRequest, resp *alihouse.TmallAlihouseTradeCouponOrderStatusQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
