package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口商品中心门店商品修改 APIRequest
alibaba.wdk.item.storesku.update

五道口商品中心门店商品修改
*/
type AlibabaWdkItemStoreskuUpdateRequest struct {
    model.Params

    // 盒马门店id
    storeId   string 

    // 商品编码
    skuCode   string 

    // 1-可售   0-不可售
    saleFlag   int64 

}

func NewAlibabaWdkItemStoreskuUpdateRequest() *AlibabaWdkItemStoreskuUpdateRequest{
    return &AlibabaWdkItemStoreskuUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkItemStoreskuUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.storesku.update"
}

func (r AlibabaWdkItemStoreskuUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkItemStoreskuUpdateRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaWdkItemStoreskuUpdateRequest) GetStoreId() string {
    return r.storeId
}

func (r *AlibabaWdkItemStoreskuUpdateRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

func (r AlibabaWdkItemStoreskuUpdateRequest) GetSkuCode() string {
    return r.skuCode
}

func (r *AlibabaWdkItemStoreskuUpdateRequest) SetSaleFlag(saleFlag int64) error {
    r.saleFlag = saleFlag
    r.Set("sale_flag", saleFlag)
    return nil
}

func (r AlibabaWdkItemStoreskuUpdateRequest) GetSaleFlag() int64 {
    return r.saleFlag
}

