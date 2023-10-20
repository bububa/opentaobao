package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrtCouponInstanceQuery 查询用户券实例
// tmall.nrt.coupon.instance.query
//
// 查询用户券实例
func TmallNrtCouponInstanceQuery(clt *core.SDKClient, req *tmallnr.TmallNrtCouponInstanceQueryAPIRequest, resp *tmallnr.TmallNrtCouponInstanceQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
