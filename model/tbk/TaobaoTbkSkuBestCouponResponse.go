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
    // Response *TaobaoTbkSkuBestCouponResponse `json:"tbk_sku_best_coupon_response,omitempty"` 
    TaobaoTbkSkuBestCouponResponse
}

/* model for simplify = false
type TaobaoTbkSkuBestCouponResponse struct {

    // 优惠券详细信息
    
    CouponInfo  *struct {
        TaobaoTbkSkuBestCouponMapData  *TaobaoTbkSkuBestCouponMapData `json:"taobao_tbk_sku_best_coupon_map_data,omitempty"`
    } `json:"coupon_info,omitempty"`
    

}
*/

type TaobaoTbkSkuBestCouponResponse struct {

    // 优惠券详细信息
    CouponInfo   *TaobaoTbkSkuBestCouponMapData `json:"coupon_info,omitempty"`

}
