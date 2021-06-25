package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据shopId和skuCode返回商品静态溯源url APIRequest
alibaba.wdk.item.trace.url.get

根据shopId和skuCode返回商品静态溯源url
*/
type AlibabaWdkItemTraceUrlGetRequest struct {
    model.Params

    // 所属门店code
    shopId   string 

    // 来源编码
    sourceCode   string 

    // 商品skuCode
    skuCode   string 

}

func NewAlibabaWdkItemTraceUrlGetRequest() *AlibabaWdkItemTraceUrlGetRequest{
    return &AlibabaWdkItemTraceUrlGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkItemTraceUrlGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.trace.url.get"
}

func (r AlibabaWdkItemTraceUrlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkItemTraceUrlGetRequest) SetShopId(shopId string) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

func (r AlibabaWdkItemTraceUrlGetRequest) GetShopId() string {
    return r.shopId
}

func (r *AlibabaWdkItemTraceUrlGetRequest) SetSourceCode(sourceCode string) error {
    r.sourceCode = sourceCode
    r.Set("source_code", sourceCode)
    return nil
}

func (r AlibabaWdkItemTraceUrlGetRequest) GetSourceCode() string {
    return r.sourceCode
}

func (r *AlibabaWdkItemTraceUrlGetRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

func (r AlibabaWdkItemTraceUrlGetRequest) GetSkuCode() string {
    return r.skuCode
}

