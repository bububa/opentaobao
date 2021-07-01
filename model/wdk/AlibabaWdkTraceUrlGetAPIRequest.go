package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
溯源url透出 API请求
alibaba.wdk.trace.url.get

根据shopId和skuCode返回商品溯源url
*/
type AlibabaWdkTraceUrlGetAPIRequest struct {
    model.Params
    // 所属门店code
    _shopId   string
    // 来源编码
    _sourceCode   string
    // barCode 或者skuCode
    _scanCode   string
}

// 初始化AlibabaWdkTraceUrlGetAPIRequest对象
func NewAlibabaWdkTraceUrlGetRequest() *AlibabaWdkTraceUrlGetAPIRequest{
    return &AlibabaWdkTraceUrlGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkTraceUrlGetAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.trace.url.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkTraceUrlGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShopId Setter
// 所属门店code
func (r *AlibabaWdkTraceUrlGetAPIRequest) SetShopId(_shopId string) error {
    r._shopId = _shopId
    r.Set("shop_id", _shopId)
    return nil
}

// ShopId Getter
func (r AlibabaWdkTraceUrlGetAPIRequest) GetShopId() string {
    return r._shopId
}
// SourceCode Setter
// 来源编码
func (r *AlibabaWdkTraceUrlGetAPIRequest) SetSourceCode(_sourceCode string) error {
    r._sourceCode = _sourceCode
    r.Set("source_code", _sourceCode)
    return nil
}

// SourceCode Getter
func (r AlibabaWdkTraceUrlGetAPIRequest) GetSourceCode() string {
    return r._sourceCode
}
// ScanCode Setter
// barCode 或者skuCode
func (r *AlibabaWdkTraceUrlGetAPIRequest) SetScanCode(_scanCode string) error {
    r._scanCode = _scanCode
    r.Set("scan_code", _scanCode)
    return nil
}

// ScanCode Getter
func (r AlibabaWdkTraceUrlGetAPIRequest) GetScanCode() string {
    return r._scanCode
}
