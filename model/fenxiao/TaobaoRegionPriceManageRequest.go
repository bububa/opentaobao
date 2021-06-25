package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑区域价格 APIRequest
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

func NewTaobaoRegionPriceManageRequest() *TaobaoRegionPriceManageRequest{
    return &TaobaoRegionPriceManageRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRegionPriceManageRequest) GetApiMethodName() string {
    return "taobao.region.price.manage"
}

func (r TaobaoRegionPriceManageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRegionPriceManageRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoRegionPriceManageRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoRegionPriceManageRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

func (r TaobaoRegionPriceManageRequest) GetSkuId() int64 {
    return r.skuId
}

func (r *TaobaoRegionPriceManageRequest) SetRegionalPriceDtos(regionalPriceDtos []RegionalPriceDto) error {
    r.regionalPriceDtos = regionalPriceDtos
    r.Set("regional_price_dtos", regionalPriceDtos)
    return nil
}

func (r TaobaoRegionPriceManageRequest) GetRegionalPriceDtos() []RegionalPriceDto {
    return r.regionalPriceDtos
}

func (r *TaobaoRegionPriceManageRequest) SetIsFull(isFull bool) error {
    r.isFull = isFull
    r.Set("is_full", isFull)
    return nil
}

func (r TaobaoRegionPriceManageRequest) GetIsFull() bool {
    return r.isFull
}

