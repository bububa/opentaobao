package tmalltrend

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmalltrend"
)

// Tmallbranditemcouponprotect 全域新品店铺优惠券免除
// tmall.brand.item.coupon.protect
//
// 全域新品店铺优惠券免除申请打标接口
func Tmallbranditemcouponprotect(clt *core.SDKClient, req *tmalltrend.TmallbranditemcouponprotectAPIRequest, session string) (*tmalltrend.TmallbranditemcouponprotectAPIResponse, error) {
	var resp tmalltrend.TmallbranditemcouponprotectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
