package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkSkuBestCouponAPIRequest
sku维度最优优惠券信息 API请求
taobao.tbk.sku.best.coupon

根据itemid和skuid查询最优优惠券信息 */
type TaobaoTbkSkuBestCouponAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
	// 商品对应的skuId
	_skuId int64
}

// New
