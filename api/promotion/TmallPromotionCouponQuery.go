package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Tmallpromotioncouponquery 查询可用优惠券列表
// tmall.promotion.coupon.query
//
// 查询用户的可用优惠券列表，仅包含优惠券基本信息和用户nick
func Tmallpromotioncouponquery(clt *core.SDKClient, req *promotion.TmallpromotioncouponqueryAPIRequest, session string) (*promotion.TmallpromotioncouponqueryAPIResponse, error) {
	var resp promotion.TmallpromotioncouponqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
