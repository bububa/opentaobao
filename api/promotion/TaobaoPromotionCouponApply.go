package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotioncouponapply 优惠券领取
// taobao.promotion.coupon.apply
//
// 优惠券领取
func Taobaopromotioncouponapply(clt *core.SDKClient, req *promotion.TaobaopromotioncouponapplyAPIRequest, session string) (*promotion.TaobaopromotioncouponapplyAPIResponse, error) {
	var resp promotion.TaobaopromotioncouponapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
