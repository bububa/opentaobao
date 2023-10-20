package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Tmallnrtcouponinstancequery 查询用户券实例
// tmall.nrt.coupon.instance.query
//
// 查询用户券实例
func Tmallnrtcouponinstancequery(clt *core.SDKClient, req *tmallnr.TmallnrtcouponinstancequeryAPIRequest, session string) (*tmallnr.TmallnrtcouponinstancequeryAPIResponse, error) {
	var resp tmallnr.TmallnrtcouponinstancequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
