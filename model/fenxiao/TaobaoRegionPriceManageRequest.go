package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑区域价格 API请求
taobao.region.price.manage

编辑区域价格
*/
type TaobaoRegionPriceManageRequest struct {
    model.Params
    // 商品id
    itemId   int64
    // 无sku传0
    skuId   int64
    // 列表
    regionalPriceDtos   []RegionalPriceDto
    // true:全量, false:增量
    isFull   bool
}

// 初始化TaobaoRegionPriceManageRequest对象
func NewTaobaoRegionPriceManageRequest() *TaobaoRegionPriceManageRequest{
    return &TaobaoRegionPriceManageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRegionPriceManageRequest) GetApiMethodName() string {
    return "taobao.region.price.manage"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRegionPriceManageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TaobaoRegionPriceManageRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoRegionPriceManageRequest) GetItemId() int64 {
    return r.itemId
}
// SkuId Setter
// 无sku传0
func (r *TaobaoRegionPriceManageRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

// SkuId Getter
func (r TaobaoRegionPriceManageRequest) GetSkuId() int64 {
    return r.skuId
}
// RegionalPriceDtos Setter
// 列表
func (r *TaobaoRegionPriceManageRequest) SetRegionalPriceDtos(regionalPriceDtos []RegionalPriceDto) error {
    r.regionalPriceDtos = regionalPriceDtos
    r.Set("regional_price_dtos", regionalPriceDtos)
    return nil
}

// RegionalPriceDtos Getter
func (r TaobaoRegionPriceManageRequest) GetRegionalPriceDtos() []RegionalPriceDto {
    return r.regionalPriceDtos
}
// IsFull Setter
// true:全量, false:增量
func (r *TaobaoRegionPriceManageRequest) SetIsFull(isFull bool) error {
    r.isFull = isFull
    r.Set("is_full", isFull)
    return nil
}

// IsFull Getter
func (r TaobaoRegionPriceManageRequest) GetIsFull() bool {
    return r.isFull
}
