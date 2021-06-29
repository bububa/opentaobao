package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
加密招商一网能支付入参 API请求
alibaba.damai.maitix.distribution.cmb.paramencrypt

encryptParam4Cmb
*/
type AlibabaDamaiMaitixDistributionCmbParamencryptRequest struct {
    model.Params
    // 入参param
    param   *DisEncrypt4CmbParam
}

// 初始化AlibabaDamaiMaitixDistributionCmbParamencryptRequest对象
func NewAlibabaDamaiMaitixDistributionCmbParamencryptRequest() *AlibabaDamaiMaitixDistributionCmbParamencryptRequest{
    return &AlibabaDamaiMaitixDistributionCmbParamencryptRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixDistributionCmbParamencryptRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.distribution.cmb.paramencrypt"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixDistributionCmbParamencryptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参param
func (r *AlibabaDamaiMaitixDistributionCmbParamencryptRequest) SetParam(param *DisEncrypt4CmbParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaDamaiMaitixDistributionCmbParamencryptRequest) GetParam() *DisEncrypt4CmbParam {
    return r.param
}
