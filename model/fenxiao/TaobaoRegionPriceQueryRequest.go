package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
区域价格查询 APIRequest
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

func NewTaobaoRegionPriceQueryRequest() *TaobaoRegionPriceQueryRequest{
    return &TaobaoRegionPriceQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRegionPriceQueryRequest) GetApiMethodName() string {
    return "taobao.region.price.query"
}

func (r TaobaoRegionPriceQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRegionPriceQueryRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoRegionPriceQueryRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoRegionPriceQueryRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

func (r TaobaoRegionPriceQueryRequest) GetSkuId() int64 {
    return r.skuId
}

func (r *TaobaoRegionPriceQueryRequest) SetRegionalPriceDtos(regionalPriceDtos []RegionalPriceDto) error {
    r.regionalPriceDtos = regionalPriceDtos
    r.Set("regional_price_dtos", regionalPriceDtos)
    return nil
}

func (r TaobaoRegionPriceQueryRequest) GetRegionalPriceDtos() []RegionalPriceDto {
    return r.regionalPriceDtos
}

