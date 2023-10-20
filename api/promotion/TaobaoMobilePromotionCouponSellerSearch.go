package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaomobilepromotioncouponsellersearch 查询绑定卖家优惠券相关信息(手淘专用)
// taobao.mobile.promotion.coupon.seller.search
//
// 查询绑定卖家相关优惠券信息 如isv 百川 等外部业务方
func Taobaomobilepromotioncouponsellersearch(clt *core.SDKClient, req *promotion.TaobaomobilepromotioncouponsellersearchAPIRequest, session string) (*promotion.TaobaomobilepromotioncouponsellersearchAPIResponse, error) {
	var resp promotion.TaobaomobilepromotioncouponsellersearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
