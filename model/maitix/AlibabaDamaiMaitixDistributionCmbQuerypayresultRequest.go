package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询招行支付状态api API请求
alibaba.damai.maitix.distribution.cmb.querypayresult

queryPayResult
*/
type AlibabaDamaiMaitixDistributionCmbQuerypayresultRequest struct {
    model.Params
    // 入参param
    _param   *QueryPayResultParam
}

// 初始化AlibabaDamaiMaitixDistributionCmbQuerypayresultRequest对象
func NewAlibabaDamaiMaitixDistributionCmbQuerypayresultRequest() *AlibabaDamaiMaitixDistributionCmbQuerypayresultRequest{
    return &AlibabaDamaiMaitixDistributionCmbQuerypayresultRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixDistributionCmbQuerypayresultRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.distribution.cmb.querypayresult"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixDistributionCmbQuerypayresultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参param
func (r *AlibabaDamaiMaitixDistributionCmbQuerypayresultRequest) SetParam(_param *QueryPayResultParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaDamaiMaitixDistributionCmbQuerypayresultRequest) GetParam() *QueryPayResultParam {
    return r._param
}
