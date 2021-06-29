package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据shopId和skuCode返回商品静态溯源url API请求
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

// 初始化AlibabaWdkItemTraceUrlGetRequest对象
func NewAlibabaWdkItemTraceUrlGetRequest() *AlibabaWdkItemTraceUrlGetRequest{
    return &AlibabaWdkItemTraceUrlGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemTraceUrlGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.trace.url.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemTraceUrlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShopId Setter
// 所属门店code
func (r *AlibabaWdkItemTraceUrlGetRequest) SetShopId(shopId string) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

// ShopId Getter
func (r AlibabaWdkItemTraceUrlGetRequest) GetShopId() string {
    return r.shopId
}
// SourceCode Setter
// 来源编码
func (r *AlibabaWdkItemTraceUrlGetRequest) SetSourceCode(sourceCode string) error {
    r.sourceCode = sourceCode
    r.Set("source_code", sourceCode)
    return nil
}

// SourceCode Getter
func (r AlibabaWdkItemTraceUrlGetRequest) GetSourceCode() string {
    return r.sourceCode
}
// SkuCode Setter
// 商品skuCode
func (r *AlibabaWdkItemTraceUrlGetRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkItemTraceUrlGetRequest) GetSkuCode() string {
    return r.skuCode
}
