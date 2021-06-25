package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
蜂鸟商户解约接口 APIRequest
alibaba.ele.fengniao.merchant.contract.cancel

通过调用此接口，商家及商家下的所有门店解除蜂鸟物流服务
*/
type AlibabaEleFengniaoMerchantContractCancelRequest struct {
    model.Params

    // 系统自动生成
    param   *Param 

}

func NewAlibabaEleFengniaoMerchantContractCancelRequest() *AlibabaEleFengniaoMerchantContractCancelRequest{
    return &AlibabaEleFengniaoMerchantContractCancelRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleFengniaoMerchantContractCancelRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.merchant.contract.cancel"
}

func (r AlibabaEleFengniaoMerchantContractCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleFengniaoMerchantContractCancelRequest) SetParam(param *Param) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaEleFengniaoMerchantContractCancelRequest) GetParam() *Param {
    return r.param
}

