package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店商品信息查询 APIRequest
alibaba.wdk.item.storesku.query

门店商品信息查询
*/
type AlibabaWdkItemStoreskuQueryRequest struct {
    model.Params

    // 商品编码
    skuCode   string 

    // 门店ID
    storeId   string 

}

func NewAlibabaWdkItemStoreskuQueryRequest() *AlibabaWdkItemStoreskuQueryRequest{
    return &AlibabaWdkItemStoreskuQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkItemStoreskuQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.storesku.query"
}

func (r AlibabaWdkItemStoreskuQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkItemStoreskuQueryRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

func (r AlibabaWdkItemStoreskuQueryRequest) GetSkuCode() string {
    return r.skuCode
}

func (r *AlibabaWdkItemStoreskuQueryRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaWdkItemStoreskuQueryRequest) GetStoreId() string {
    return r.storeId
}

