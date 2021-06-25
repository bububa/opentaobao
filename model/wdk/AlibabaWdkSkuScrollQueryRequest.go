package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店商品批量游标方式查询接口 APIRequest
alibaba.wdk.sku.scroll.query

通过游标方式批量获取门店商品信息，包括商品条码，商品名称，价格，会员价等信息。
*/
type AlibabaWdkSkuScrollQueryRequest struct {
    model.Params

    // 商家类目编码
    merchantCatCode   string 

    // 门店编码
    ouCode   string 

    // 游标：第一次请求不用填写，否则请填写上一次请求返回的值，直到获取到足够的数据
    scrollId   string 

    // 英文逗号分隔的商品编码，最多20个。如果配合门店字段使用，直接非游标方式返回商品数据
    skuCodes   string 

}

func NewAlibabaWdkSkuScrollQueryRequest() *AlibabaWdkSkuScrollQueryRequest{
    return &AlibabaWdkSkuScrollQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSkuScrollQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.scroll.query"
}

func (r AlibabaWdkSkuScrollQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSkuScrollQueryRequest) SetMerchantCatCode(merchantCatCode string) error {
    r.merchantCatCode = merchantCatCode
    r.Set("merchant_cat_code", merchantCatCode)
    return nil
}

func (r AlibabaWdkSkuScrollQueryRequest) GetMerchantCatCode() string {
    return r.merchantCatCode
}

func (r *AlibabaWdkSkuScrollQueryRequest) SetOuCode(ouCode string) error {
    r.ouCode = ouCode
    r.Set("ou_code", ouCode)
    return nil
}

func (r AlibabaWdkSkuScrollQueryRequest) GetOuCode() string {
    return r.ouCode
}

func (r *AlibabaWdkSkuScrollQueryRequest) SetScrollId(scrollId string) error {
    r.scrollId = scrollId
    r.Set("scroll_id", scrollId)
    return nil
}

func (r AlibabaWdkSkuScrollQueryRequest) GetScrollId() string {
    return r.scrollId
}

func (r *AlibabaWdkSkuScrollQueryRequest) SetSkuCodes(skuCodes string) error {
    r.skuCodes = skuCodes
    r.Set("sku_codes", skuCodes)
    return nil
}

func (r AlibabaWdkSkuScrollQueryRequest) GetSkuCodes() string {
    return r.skuCodes
}

