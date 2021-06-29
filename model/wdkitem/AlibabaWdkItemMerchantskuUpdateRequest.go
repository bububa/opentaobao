package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家商品修改 APIRequest
alibaba.wdk.item.merchantsku.update

商家商品修改
*/
type AlibabaWdkItemMerchantskuUpdateRequest struct {
    model.Params

    // 商品编码
    skuCode   string 

    // 参数json
    params   string 

    // 机构编码
    orgCode   string 

}

func NewAlibabaWdkItemMerchantskuUpdateRequest() *AlibabaWdkItemMerchantskuUpdateRequest{
    return &AlibabaWdkItemMerchantskuUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkItemMerchantskuUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.merchantsku.update"
}

func (r AlibabaWdkItemMerchantskuUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkItemMerchantskuUpdateRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

func (r AlibabaWdkItemMerchantskuUpdateRequest) GetSkuCode() string {
    return r.skuCode
}

func (r *AlibabaWdkItemMerchantskuUpdateRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

func (r AlibabaWdkItemMerchantskuUpdateRequest) GetParams() string {
    return r.params
}

func (r *AlibabaWdkItemMerchantskuUpdateRequest) SetOrgCode(orgCode string) error {
    r.orgCode = orgCode
    r.Set("org_code", orgCode)
    return nil
}

func (r AlibabaWdkItemMerchantskuUpdateRequest) GetOrgCode() string {
    return r.orgCode
}

