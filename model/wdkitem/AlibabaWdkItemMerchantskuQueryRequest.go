package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家商品信息查询 APIRequest
alibaba.wdk.item.merchantsku.query

商家商品信息查询
*/
type AlibabaWdkItemMerchantskuQueryRequest struct {
    model.Params

    // 商品编码
    skuCode   string 

    // 机构编码
    orgCode   string 

}

func NewAlibabaWdkItemMerchantskuQueryRequest() *AlibabaWdkItemMerchantskuQueryRequest{
    return &AlibabaWdkItemMerchantskuQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkItemMerchantskuQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.merchantsku.query"
}

func (r AlibabaWdkItemMerchantskuQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkItemMerchantskuQueryRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

func (r AlibabaWdkItemMerchantskuQueryRequest) GetSkuCode() string {
    return r.skuCode
}

func (r *AlibabaWdkItemMerchantskuQueryRequest) SetOrgCode(orgCode string) error {
    r.orgCode = orgCode
    r.Set("org_code", orgCode)
    return nil
}

func (r AlibabaWdkItemMerchantskuQueryRequest) GetOrgCode() string {
    return r.orgCode
}

