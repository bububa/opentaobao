package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（供销）产品级别日历价格库存修改，全量覆盖 API请求
taobao.alitrip.travel.product.sku.override

（供销）产品级别日历价格库存修改，全量覆盖
*/
type TaobaoAlitripTravelProductSkuOverrideRequest struct {
    model.Params
    // 商品 外部商家编码。itemId和outProductId至少填写一个
    _outProductId   string
    // 商品id。itemId和outProductId至少填写一个
    _itemId   int64
    // 商品日历价格库存套餐
    _skus   []ItemSkuInfo
}

// 初始化TaobaoAlitripTravelProductSkuOverrideRequest对象
func NewTaobaoAlitripTravelProductSkuOverrideRequest() *TaobaoAlitripTravelProductSkuOverrideRequest{
    return &TaobaoAlitripTravelProductSkuOverrideRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelProductSkuOverrideRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.product.sku.override"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelProductSkuOverrideRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutProductId Setter
// 商品 外部商家编码。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelProductSkuOverrideRequest) SetOutProductId(_outProductId string) error {
    r._outProductId = _outProductId
    r.Set("out_product_id", _outProductId)
    return nil
}

// OutProductId Getter
func (r TaobaoAlitripTravelProductSkuOverrideRequest) GetOutProductId() string {
    return r._outProductId
}
// ItemId Setter
// 商品id。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelProductSkuOverrideRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoAlitripTravelProductSkuOverrideRequest) GetItemId() int64 {
    return r._itemId
}
// Skus Setter
// 商品日历价格库存套餐
func (r *TaobaoAlitripTravelProductSkuOverrideRequest) SetSkus(_skus []ItemSkuInfo) error {
    r._skus = _skus
    r.Set("skus", _skus)
    return nil
}

// Skus Getter
func (r TaobaoAlitripTravelProductSkuOverrideRequest) GetSkus() []ItemSkuInfo {
    return r._skus
}
