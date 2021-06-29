package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家产品服务-编辑产品 API请求
alibaba.wdk.merchantproduct.edit

商家产品服务-编辑产品
*/
type AlibabaWdkMerchantproductEditRequest struct {
    model.Params
    // 产品编辑入参
    _req   *MerchantProductRequest
}

// 初始化AlibabaWdkMerchantproductEditRequest对象
func NewAlibabaWdkMerchantproductEditRequest() *AlibabaWdkMerchantproductEditRequest{
    return &AlibabaWdkMerchantproductEditRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantproductEditRequest) GetApiMethodName() string {
    return "alibaba.wdk.merchantproduct.edit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantproductEditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Req Setter
// 产品编辑入参
func (r *AlibabaWdkMerchantproductEditRequest) SetReq(_req *MerchantProductRequest) error {
    r._req = _req
    r.Set("req", _req)
    return nil
}

// Req Getter
func (r AlibabaWdkMerchantproductEditRequest) GetReq() *MerchantProductRequest {
    return r._req
}
