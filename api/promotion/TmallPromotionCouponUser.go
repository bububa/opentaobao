package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TmallPromotionCouponUser 用户信息查询接口
// tmall.promotion.coupon.user
//
// 开发给外部合作商（例如：苏宁），通过会员付款码获得会员nick
func TmallPromotionCouponUser(clt *core.SDKClient, req *promotion.TmallPromotionCouponUserAPIRequest, session string) (*promotion.TmallPromotionCouponUserAPIResponse, error) {
	var resp promotion.TmallPromotionCouponUserAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
