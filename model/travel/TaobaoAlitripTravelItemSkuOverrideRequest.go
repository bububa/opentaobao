package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】商品级别日历价格库存修改，全量覆盖 API请求
taobao.alitrip.travel.item.sku.override

旅行度假新商品日历价格库存信息修改接口 第三版。提供商家通过TOP API方式修改商品sku信息。
*/
type TaobaoAlitripTravelItemSkuOverrideRequest struct {
    model.Params
    // 商品id。itemId和outProductId至少填写一个
    itemId   int64
    // 商品日历价格库存套餐
    skus   []PontusTravelItemSkuInfo
    // 商品 外部商家编码。itemId和outProductId至少填写一个
    outProductId   string
}

// 初始化TaobaoAlitripTravelItemSkuOverrideRequest对象
func NewTaobaoAlitripTravelItemSkuOverrideRequest() *TaobaoAlitripTravelItemSkuOverrideRequest{
    return &TaobaoAlitripTravelItemSkuOverrideRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelItemSkuOverrideRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.item.sku.override"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelItemSkuOverrideRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemSkuOverrideRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoAlitripTravelItemSkuOverrideRequest) GetItemId() int64 {
    return r.itemId
}
// Skus Setter
// 商品日历价格库存套餐
func (r *TaobaoAlitripTravelItemSkuOverrideRequest) SetSkus(skus []PontusTravelItemSkuInfo) error {
    r.skus = skus
    r.Set("skus", skus)
    return nil
}

// Skus Getter
func (r TaobaoAlitripTravelItemSkuOverrideRequest) GetSkus() []PontusTravelItemSkuInfo {
    return r.skus
}
// OutProductId Setter
// 商品 外部商家编码。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemSkuOverrideRequest) SetOutProductId(outProductId string) error {
    r.outProductId = outProductId
    r.Set("out_product_id", outProductId)
    return nil
}

// OutProductId Getter
func (r TaobaoAlitripTravelItemSkuOverrideRequest) GetOutProductId() string {
    return r.outProductId
}
