package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkSkuBestCoupon sku维度最优优惠券信息
// taobao.tbk.sku.best.coupon
//
// 根据itemid和skuid查询最优优惠券信息
func TaobaoTbkSkuBestCoupon(clt *core.SDKClient, req *tbk.TaobaoTbkSkuBestCouponAPIRequest, resp *tbk.TaobaoTbkSkuBestCouponAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
