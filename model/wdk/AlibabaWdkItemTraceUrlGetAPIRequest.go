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
type AlibabaWdkItemTraceUrlGetAPIRequest struct {
    model.Params
    // 所属门店code
    _shopId   string
    // 来源编码
    _sourceCode   string
    // 商品skuCode
    _skuCode   string
}

// 初始化AlibabaWdkItemTraceUrlGetAPIRequest对象
func NewAlibabaWdkItemTraceUrlGetRequest() *AlibabaWdkItemTraceUrlGetAPIRequest{
    return &AlibabaWdkItemTraceUrlGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemTraceUrlGetAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.trace.url.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemTraceUrlGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShopId Setter
// 所属门店code
func (r *AlibabaWdkItemTraceUrlGetAPIRequest) SetShopId(_shopId string) error {
    r._shopId = _shopId
    r.Set("shop_id", _shopId)
    return nil
}

// ShopId Getter
func (r AlibabaWdkItemTraceUrlGetAPIRequest) GetShopId() string {
    return r._shopId
}
// SourceCode Setter
// 来源编码
func (r *AlibabaWdkItemTraceUrlGetAPIRequest) SetSourceCode(_sourceCode string) error {
    r._sourceCode = _sourceCode
    r.Set("source_code", _sourceCode)
    return nil
}

// SourceCode Getter
func (r AlibabaWdkItemTraceUrlGetAPIRequest) GetSourceCode() string {
    return r._sourceCode
}
// SkuCode Setter
// 商品skuCode
func (r *AlibabaWdkItemTraceUrlGetAPIRequest) SetSkuCode(_skuCode string) error {
    r._skuCode = _skuCode
    r.Set("sku_code", _skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkItemTraceUrlGetAPIRequest) GetSkuCode() string {
    return r._skuCode
}
