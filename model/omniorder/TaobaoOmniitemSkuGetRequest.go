package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取全渠道门店商品sku API请求
taobao.omniitem.sku.get

通过skuId或者skuOutId查询全渠道门店商品sku信息
*/
type TaobaoOmniitemSkuGetRequest struct {
    model.Params
    // 商品id
    _itemId   int64
    // skuId
    _skuId   int64
    // sku商家编码
    _skuOuterId   string
}

// 初始化TaobaoOmniitemSkuGetRequest对象
func NewTaobaoOmniitemSkuGetRequest() *TaobaoOmniitemSkuGetRequest{
    return &TaobaoOmniitemSkuGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemSkuGetRequest) GetApiMethodName() string {
    return "taobao.omniitem.sku.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemSkuGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TaobaoOmniitemSkuGetRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOmniitemSkuGetRequest) GetItemId() int64 {
    return r._itemId
}
// SkuId Setter
// skuId
func (r *TaobaoOmniitemSkuGetRequest) SetSkuId(_skuId int64) error {
    r._skuId = _skuId
    r.Set("sku_id", _skuId)
    return nil
}

// SkuId Getter
func (r TaobaoOmniitemSkuGetRequest) GetSkuId() int64 {
    return r._skuId
}
// SkuOuterId Setter
// sku商家编码
func (r *TaobaoOmniitemSkuGetRequest) SetSkuOuterId(_skuOuterId string) error {
    r._skuOuterId = _skuOuterId
    r.Set("sku_outer_id", _skuOuterId)
    return nil
}

// SkuOuterId Getter
func (r TaobaoOmniitemSkuGetRequest) GetSkuOuterId() string {
    return r._skuOuterId
}
