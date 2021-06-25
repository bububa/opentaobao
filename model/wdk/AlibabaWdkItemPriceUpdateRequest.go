package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口商品中心价格修改 APIRequest
alibaba.wdk.item.price.update

修改门店商品售价和会员价
*/
type AlibabaWdkItemPriceUpdateRequest struct {
    model.Params

    // 盒马门店id
    storeId   string 

    // 商品编码
    skuCode   string 

    // 价格，单位是分
    skuPrice   int64 

    // 会员价格，单位是分，且不能大于价格
    skuMemberPrice   int64 

}

func NewAlibabaWdkItemPriceUpdateRequest() *AlibabaWdkItemPriceUpdateRequest{
    return &AlibabaWdkItemPriceUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkItemPriceUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.price.update"
}

func (r AlibabaWdkItemPriceUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkItemPriceUpdateRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaWdkItemPriceUpdateRequest) GetStoreId() string {
    return r.storeId
}

func (r *AlibabaWdkItemPriceUpdateRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

func (r AlibabaWdkItemPriceUpdateRequest) GetSkuCode() string {
    return r.skuCode
}

func (r *AlibabaWdkItemPriceUpdateRequest) SetSkuPrice(skuPrice int64) error {
    r.skuPrice = skuPrice
    r.Set("sku_price", skuPrice)
    return nil
}

func (r AlibabaWdkItemPriceUpdateRequest) GetSkuPrice() int64 {
    return r.skuPrice
}

func (r *AlibabaWdkItemPriceUpdateRequest) SetSkuMemberPrice(skuMemberPrice int64) error {
    r.skuMemberPrice = skuMemberPrice
    r.Set("sku_member_price", skuMemberPrice)
    return nil
}

func (r AlibabaWdkItemPriceUpdateRequest) GetSkuMemberPrice() int64 {
    return r.skuMemberPrice
}

