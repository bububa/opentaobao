package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
区域价格查询 API请求
taobao.region.price.query

区域价格查询
*/
type TaobaoRegionPriceQueryRequest struct {
    model.Params
    // 商品id
    itemId   int64
    // 无sku可传0
    skuId   int64
    // 不传则返回所有设置的区域价格
    regionalPriceDtos   []RegionalPriceDto
}

// 初始化TaobaoRegionPriceQueryRequest对象
func NewTaobaoRegionPriceQueryRequest() *TaobaoRegionPriceQueryRequest{
    return &TaobaoRegionPriceQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRegionPriceQueryRequest) GetApiMethodName() string {
    return "taobao.region.price.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRegionPriceQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TaobaoRegionPriceQueryRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoRegionPriceQueryRequest) GetItemId() int64 {
    return r.itemId
}
// SkuId Setter
// 无sku可传0
func (r *TaobaoRegionPriceQueryRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

// SkuId Getter
func (r TaobaoRegionPriceQueryRequest) GetSkuId() int64 {
    return r.skuId
}
// RegionalPriceDtos Setter
// 不传则返回所有设置的区域价格
func (r *TaobaoRegionPriceQueryRequest) SetRegionalPriceDtos(regionalPriceDtos []RegionalPriceDto) error {
    r.regionalPriceDtos = regionalPriceDtos
    r.Set("regional_price_dtos", regionalPriceDtos)
    return nil
}

// RegionalPriceDtos Getter
func (r TaobaoRegionPriceQueryRequest) GetRegionalPriceDtos() []RegionalPriceDto {
    return r.regionalPriceDtos
}
