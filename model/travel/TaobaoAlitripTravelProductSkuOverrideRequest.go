package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（供销）产品级别日历价格库存修改，全量覆盖 APIRequest
taobao.alitrip.travel.product.sku.override

（供销）产品级别日历价格库存修改，全量覆盖
*/
type TaobaoAlitripTravelProductSkuOverrideRequest struct {
    model.Params

    // 商品 外部商家编码。itemId和outProductId至少填写一个
    outProductId   string 

    // 商品id。itemId和outProductId至少填写一个
    itemId   int64 

    // 商品日历价格库存套餐
    skus   []ItemSkuInfo 

}

func NewTaobaoAlitripTravelProductSkuOverrideRequest() *TaobaoAlitripTravelProductSkuOverrideRequest{
    return &TaobaoAlitripTravelProductSkuOverrideRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripTravelProductSkuOverrideRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.product.sku.override"
}

func (r TaobaoAlitripTravelProductSkuOverrideRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripTravelProductSkuOverrideRequest) SetOutProductId(outProductId string) error {
    r.outProductId = outProductId
    r.Set("out_product_id", outProductId)
    return nil
}

func (r TaobaoAlitripTravelProductSkuOverrideRequest) GetOutProductId() string {
    return r.outProductId
}

func (r *TaobaoAlitripTravelProductSkuOverrideRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoAlitripTravelProductSkuOverrideRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoAlitripTravelProductSkuOverrideRequest) SetSkus(skus []ItemSkuInfo) error {
    r.skus = skus
    r.Set("skus", skus)
    return nil
}

func (r TaobaoAlitripTravelProductSkuOverrideRequest) GetSkus() []ItemSkuInfo {
    return r.skus
}

