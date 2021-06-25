package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家产品服务-编辑产品 APIRequest
alibaba.wdk.merchantproduct.edit

商家产品服务-编辑产品
*/
type AlibabaWdkMerchantproductEditRequest struct {
    model.Params

    // 产品编辑入参
    req   *MerchantProductRequest 

}

func NewAlibabaWdkMerchantproductEditRequest() *AlibabaWdkMerchantproductEditRequest{
    return &AlibabaWdkMerchantproductEditRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMerchantproductEditRequest) GetApiMethodName() string {
    return "alibaba.wdk.merchantproduct.edit"
}

func (r AlibabaWdkMerchantproductEditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMerchantproductEditRequest) SetReq(req *MerchantProductRequest) error {
    r.req = req
    r.Set("req", req)
    return nil
}

func (r AlibabaWdkMerchantproductEditRequest) GetReq() *MerchantProductRequest {
    return r.req
}

