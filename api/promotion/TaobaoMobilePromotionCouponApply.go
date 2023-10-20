package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaomobilepromotioncouponapply 优惠券领取(手淘专用)
// taobao.mobile.promotion.coupon.apply
//
// 优惠券领取
func Taobaomobilepromotioncouponapply(clt *core.SDKClient, req *promotion.TaobaomobilepromotioncouponapplyAPIRequest, session string) (*promotion.TaobaomobilepromotioncouponapplyAPIResponse, error) {
	var resp promotion.TaobaomobilepromotioncouponapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
