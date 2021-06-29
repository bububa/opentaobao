package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消区域价格 API请求
taobao.region.price.cancle

取消区域价格
*/
type TaobaoRegionPriceCancleRequest struct {
    model.Params
    // 商品
    _itemId   int64
    // 无sku传0
    _skuId   int64
}

// 初始化TaobaoRegionPriceCancleRequest对象
func NewTaobaoRegionPriceCancleRequest() *TaobaoRegionPriceCancleRequest{
    return &TaobaoRegionPriceCancleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRegionPriceCancleRequest) GetApiMethodName() string {
    return "taobao.region.price.cancle"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRegionPriceCancleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品
func (r *TaobaoRegionPriceCancleRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoRegionPriceCancleRequest) GetItemId() int64 {
    return r._itemId
}
// SkuId Setter
// 无sku传0
func (r *TaobaoRegionPriceCancleRequest) SetSkuId(_skuId int64) error {
    r._skuId = _skuId
    r.Set("sku_id", _skuId)
    return nil
}

// SkuId Getter
func (r TaobaoRegionPriceCancleRequest) GetSkuId() int64 {
    return r._skuId
}
