package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
sku维度最优优惠券信息 APIRequest
taobao.tbk.sku.best.coupon

根据itemid和skuid查询最优优惠券信息
*/
type TaobaoTbkSkuBestCouponRequest struct {
    model.Params

    // 商品ID
    itemId   int64 

    // 商品对应的skuId
    skuId   int64 

}

func NewTaobaoTbkSkuBestCouponRequest() *TaobaoTbkSkuBestCouponRequest{
    return &TaobaoTbkSkuBestCouponRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTbkSkuBestCouponRequest) GetApiMethodName() string {
    return "taobao.tbk.sku.best.coupon"
}

func (r TaobaoTbkSkuBestCouponRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTbkSkuBestCouponRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoTbkSkuBestCouponRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoTbkSkuBestCouponRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

func (r TaobaoTbkSkuBestCouponRequest) GetSkuId() int64 {
    return r.skuId
}

