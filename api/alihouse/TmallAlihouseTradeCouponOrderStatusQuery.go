package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Tmallalihousetradecouponorderstatusquery 查询电商券履约单状态
// tmall.alihouse.trade.coupon.order.status.query
//
// 查询电商券履约单状态
func Tmallalihousetradecouponorderstatusquery(clt *core.SDKClient, req *alihouse.TmallalihousetradecouponorderstatusqueryAPIRequest, session string) (*alihouse.TmallalihousetradecouponorderstatusqueryAPIResponse, error) {
	var resp alihouse.TmallalihousetradecouponorderstatusqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
