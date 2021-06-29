package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】套餐级别日历价格库存增删操作 API请求
taobao.alitrip.travel.item.sku.package.modify

【API3.0】套餐级别日历价格库存增删操作
*/
type TaobaoAlitripTravelItemSkuPackageModifyRequest struct {
    model.Params
    // 商品id。itemId和outProductId至少填写一个
    _itemId   int64
    // 商品 外部商家编码。itemId和outProductId至少填写一个
    _outProductId   string
    // 商品日历价格库存套餐
    _skus   []ItemSkuInfo
}

// 初始化TaobaoAlitripTravelItemSkuPackageModifyRequest对象
func NewTaobaoAlitripTravelItemSkuPackageModifyRequest() *TaobaoAlitripTravelItemSkuPackageModifyRequest{
    return &TaobaoAlitripTravelItemSkuPackageModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelItemSkuPackageModifyRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.item.sku.package.modify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelItemSkuPackageModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemSkuPackageModifyRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoAlitripTravelItemSkuPackageModifyRequest) GetItemId() int64 {
    return r._itemId
}
// OutProductId Setter
// 商品 外部商家编码。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemSkuPackageModifyRequest) SetOutProductId(_outProductId string) error {
    r._outProductId = _outProductId
    r.Set("out_product_id", _outProductId)
    return nil
}

// OutProductId Getter
func (r TaobaoAlitripTravelItemSkuPackageModifyRequest) GetOutProductId() string {
    return r._outProductId
}
// Skus Setter
// 商品日历价格库存套餐
func (r *TaobaoAlitripTravelItemSkuPackageModifyRequest) SetSkus(_skus []ItemSkuInfo) error {
    r._skus = _skus
    r.Set("skus", _skus)
    return nil
}

// Skus Getter
func (r TaobaoAlitripTravelItemSkuPackageModifyRequest) GetSkus() []ItemSkuInfo {
    return r._skus
}
