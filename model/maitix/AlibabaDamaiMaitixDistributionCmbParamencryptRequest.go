package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
加密招商一网能支付入参 APIRequest
alibaba.damai.maitix.distribution.cmb.paramencrypt

encryptParam4Cmb
*/
type AlibabaDamaiMaitixDistributionCmbParamencryptRequest struct {
    model.Params

    // 入参param
    param   *DisEncrypt4CmbParam 

}

func NewAlibabaDamaiMaitixDistributionCmbParamencryptRequest() *AlibabaDamaiMaitixDistributionCmbParamencryptRequest{
    return &AlibabaDamaiMaitixDistributionCmbParamencryptRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMaitixDistributionCmbParamencryptRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.distribution.cmb.paramencrypt"
}

func (r AlibabaDamaiMaitixDistributionCmbParamencryptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMaitixDistributionCmbParamencryptRequest) SetParam(param *DisEncrypt4CmbParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaDamaiMaitixDistributionCmbParamencryptRequest) GetParam() *DisEncrypt4CmbParam {
    return r.param
}

