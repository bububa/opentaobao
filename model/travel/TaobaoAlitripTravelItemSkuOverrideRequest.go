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
    _itemId   int64
    // 商品日历价格库存套餐
    _skus   []PontusTravelItemSkuInfo
    // 商品 外部商家编码。itemId和outProductId至少填写一个
    _outProductId   string
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
func (r *TaobaoAlitripTravelItemSkuOverrideRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoAlitripTravelItemSkuOverrideRequest) GetItemId() int64 {
    return r._itemId
}
// Skus Setter
// 商品日历价格库存套餐
func (r *TaobaoAlitripTravelItemSkuOverrideRequest) SetSkus(_skus []PontusTravelItemSkuInfo) error {
    r._skus = _skus
    r.Set("skus", _skus)
    return nil
}

// Skus Getter
func (r TaobaoAlitripTravelItemSkuOverrideRequest) GetSkus() []PontusTravelItemSkuInfo {
    return r._skus
}
// OutProductId Setter
// 商品 外部商家编码。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemSkuOverrideRequest) SetOutProductId(_outProductId string) error {
    r._outProductId = _outProductId
    r.Set("out_product_id", _outProductId)
    return nil
}

// OutProductId Getter
func (r TaobaoAlitripTravelItemSkuOverrideRequest) GetOutProductId() string {
    return r._outProductId
}
