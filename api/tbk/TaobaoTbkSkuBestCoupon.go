package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkskubestcoupon sku维度最优优惠券信息
// taobao.tbk.sku.best.coupon
//
// 根据itemid和skuid查询最优优惠券信息
func Taobaotbkskubestcoupon(clt *core.SDKClient, req *tbk.TaobaotbkskubestcouponAPIRequest, session string) (*tbk.TaobaotbkskubestcouponAPIResponse, error) {
	var resp tbk.TaobaotbkskubestcouponAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
