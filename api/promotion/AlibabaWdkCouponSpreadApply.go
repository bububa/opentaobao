package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Alibabawdkcouponspreadapply 普通发券
// alibaba.wdk.coupon.spread.apply
//
// 优惠券发放
func Alibabawdkcouponspreadapply(clt *core.SDKClient, req *promotion.AlibabawdkcouponspreadapplyAPIRequest, session string) (*promotion.AlibabawdkcouponspreadapplyAPIResponse, error) {
	var resp promotion.AlibabawdkcouponspreadapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
