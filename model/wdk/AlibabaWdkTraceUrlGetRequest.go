package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
溯源url透出 APIRequest
alibaba.wdk.trace.url.get

根据shopId和skuCode返回商品溯源url
*/
type AlibabaWdkTraceUrlGetRequest struct {
    model.Params

    // 所属门店code
    shopId   string 

    // 来源编码
    sourceCode   string 

    // barCode 或者skuCode
    scanCode   string 

}

func NewAlibabaWdkTraceUrlGetRequest() *AlibabaWdkTraceUrlGetRequest{
    return &AlibabaWdkTraceUrlGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkTraceUrlGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.trace.url.get"
}

func (r AlibabaWdkTraceUrlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkTraceUrlGetRequest) SetShopId(shopId string) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

func (r AlibabaWdkTraceUrlGetRequest) GetShopId() string {
    return r.shopId
}

func (r *AlibabaWdkTraceUrlGetRequest) SetSourceCode(sourceCode string) error {
    r.sourceCode = sourceCode
    r.Set("source_code", sourceCode)
    return nil
}

func (r AlibabaWdkTraceUrlGetRequest) GetSourceCode() string {
    return r.sourceCode
}

func (r *AlibabaWdkTraceUrlGetRequest) SetScanCode(scanCode string) error {
    r.scanCode = scanCode
    r.Set("scan_code", scanCode)
    return nil
}

func (r AlibabaWdkTraceUrlGetRequest) GetScanCode() string {
    return r.scanCode
}

