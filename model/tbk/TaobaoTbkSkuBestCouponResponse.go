package tbk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
sku维度最优优惠券信息 APIResponse
taobao.tbk.sku.best.coupon

根据itemid和skuid查询最优优惠券信息
*/
type TaobaoTbkSkuBestCouponAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTbkSkuBestCouponResponse `json:"taobao_tbk_sku_best_coupon_response,omitempty"`
}

type TaobaoTbkSkuBestCouponResponse struct {

    // 优惠券详细信息
    CouponInfo   *MapData `json:"coupon_info,omitempty"`

}
