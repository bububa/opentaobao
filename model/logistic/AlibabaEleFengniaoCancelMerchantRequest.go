package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商户取消 APIRequest
alibaba.ele.fengniao.cancel.merchant

商户取消配送
*/
type AlibabaEleFengniaoCancelMerchantRequest struct {
    model.Params

    // 参数param
    param   *Param 

}

func NewAlibabaEleFengniaoCancelMerchantRequest() *AlibabaEleFengniaoCancelMerchantRequest{
    return &AlibabaEleFengniaoCancelMerchantRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleFengniaoCancelMerchantRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.cancel.merchant"
}

func (r AlibabaEleFengniaoCancelMerchantRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleFengniaoCancelMerchantRequest) SetParam(param *Param) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaEleFengniaoCancelMerchantRequest) GetParam() *Param {
    return r.param
}

