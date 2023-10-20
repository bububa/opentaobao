package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotioncouponsellersearch 查询绑定卖家优惠券相关信息
// taobao.promotion.coupon.seller.search
//
// 查询绑定卖家相关优惠券信息  如isv  百川 等外部业务方
func Taobaopromotioncouponsellersearch(clt *core.SDKClient, req *promotion.TaobaopromotioncouponsellersearchAPIRequest, session string) (*promotion.TaobaopromotioncouponsellersearchAPIResponse, error) {
	var resp promotion.TaobaopromotioncouponsellersearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
