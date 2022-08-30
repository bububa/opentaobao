package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrtCouponInstanceQuery 查询用户券实例
// tmall.nrt.coupon.instance.query
//
// 查询用户券实例
func TmallNrtCouponInstanceQuery(clt *core.SDKClient, req *tmallnr.TmallNrtCouponInstanceQueryAPIRequest, session string) (*tmallnr.TmallNrtCouponInstanceQueryAPIResponse, error) {
	var resp tmallnr.TmallNrtCouponInstanceQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
