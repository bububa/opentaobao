package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取全渠道门店商品sku APIRequest
taobao.omniitem.sku.get

通过skuId或者skuOutId查询全渠道门店商品sku信息
*/
type TaobaoOmniitemSkuGetRequest struct {
    model.Params

    // 商品id
    itemId   int64 

    // skuId
    skuId   int64 

    // sku商家编码
    skuOuterId   string 

}

func NewTaobaoOmniitemSkuGetRequest() *TaobaoOmniitemSkuGetRequest{
    return &TaobaoOmniitemSkuGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniitemSkuGetRequest) GetApiMethodName() string {
    return "taobao.omniitem.sku.get"
}

func (r TaobaoOmniitemSkuGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniitemSkuGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoOmniitemSkuGetRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoOmniitemSkuGetRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

func (r TaobaoOmniitemSkuGetRequest) GetSkuId() int64 {
    return r.skuId
}

func (r *TaobaoOmniitemSkuGetRequest) SetSkuOuterId(skuOuterId string) error {
    r.skuOuterId = skuOuterId
    r.Set("sku_outer_id", skuOuterId)
    return nil
}

func (r TaobaoOmniitemSkuGetRequest) GetSkuOuterId() string {
    return r.skuOuterId
}

