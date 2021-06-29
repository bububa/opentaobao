package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据商品id获取商品属性 APIRequest
alibaba.alihealth.tracecodeseller.product.attr.search

根据商品id获取商品属性
*/
type AlibabaAlihealthTracecodesellerProductAttrSearchRequest struct {
    model.Params

    // 企业id
    entInfoId   int64 

    // 货品id
    tracUserProductInfoId   int64 

}

func NewAlibabaAlihealthTracecodesellerProductAttrSearchRequest() *AlibabaAlihealthTracecodesellerProductAttrSearchRequest{
    return &AlibabaAlihealthTracecodesellerProductAttrSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthTracecodesellerProductAttrSearchRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.product.attr.search"
}

func (r AlibabaAlihealthTracecodesellerProductAttrSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthTracecodesellerProductAttrSearchRequest) SetEntInfoId(entInfoId int64) error {
    r.entInfoId = entInfoId
    r.Set("ent_info_id", entInfoId)
    return nil
}

func (r AlibabaAlihealthTracecodesellerProductAttrSearchRequest) GetEntInfoId() int64 {
    return r.entInfoId
}

func (r *AlibabaAlihealthTracecodesellerProductAttrSearchRequest) SetTracUserProductInfoId(tracUserProductInfoId int64) error {
    r.tracUserProductInfoId = tracUserProductInfoId
    r.Set("trac_user_product_info_id", tracUserProductInfoId)
    return nil
}

func (r AlibabaAlihealthTracecodesellerProductAttrSearchRequest) GetTracUserProductInfoId() int64 {
    return r.tracUserProductInfoId
}

