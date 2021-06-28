package tbk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
sku维度最优优惠券信息 APIResponse
taobao.tbk.sku.best.coupon

根据itemid和skuid查询最优优惠券信息
*/
type TaobaoTbkSkuBestCouponAPIResponse struct {
    model.CommonResponse
    TaobaoTbkSkuBestCouponResponse
}

type TaobaoTbkSkuBestCouponResponse struct {
    XMLName xml.Name `xml:"tbk_sku_best_coupon_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 优惠券详细信息
    
    CouponInfo   *TaobaoTbkSkuBestCouponMapData `json:"coupon_info,omitempty" xml:"coupon_info,omitempty"`

    
}
