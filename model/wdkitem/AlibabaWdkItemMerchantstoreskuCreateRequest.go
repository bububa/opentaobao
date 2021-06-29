package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店商品信息新建 APIRequest
alibaba.wdk.item.merchantstoresku.create

门店商品信息新建
*/
type AlibabaWdkItemMerchantstoreskuCreateRequest struct {
    model.Params

    // 门店编码
    storeId   string 

    // 商品编码
    skuCode   string 

    // 新建参数，是个json字符串
    params   string 

    // 机构
    orgCode   string 

}

func NewAlibabaWdkItemMerchantstoreskuCreateRequest() *AlibabaWdkItemMerchantstoreskuCreateRequest{
    return &AlibabaWdkItemMerchantstoreskuCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkItemMerchantstoreskuCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.merchantstoresku.create"
}

func (r AlibabaWdkItemMerchantstoreskuCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkItemMerchantstoreskuCreateRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaWdkItemMerchantstoreskuCreateRequest) GetStoreId() string {
    return r.storeId
}

func (r *AlibabaWdkItemMerchantstoreskuCreateRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

func (r AlibabaWdkItemMerchantstoreskuCreateRequest) GetSkuCode() string {
    return r.skuCode
}

func (r *AlibabaWdkItemMerchantstoreskuCreateRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

func (r AlibabaWdkItemMerchantstoreskuCreateRequest) GetParams() string {
    return r.params
}

func (r *AlibabaWdkItemMerchantstoreskuCreateRequest) SetOrgCode(orgCode string) error {
    r.orgCode = orgCode
    r.Set("org_code", orgCode)
    return nil
}

func (r AlibabaWdkItemMerchantstoreskuCreateRequest) GetOrgCode() string {
    return r.orgCode
}

