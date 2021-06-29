package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销电子票查询接口 API请求
alibaba.damai.maitix.eticket.distribution.query

分销电子票查询接口
*/
type AlibabaDamaiMaitixEticketDistributionQueryRequest struct {
    model.Params
    // 入参param
    param   *EticketQueryParam
}

// 初始化AlibabaDamaiMaitixEticketDistributionQueryRequest对象
func NewAlibabaDamaiMaitixEticketDistributionQueryRequest() *AlibabaDamaiMaitixEticketDistributionQueryRequest{
    return &AlibabaDamaiMaitixEticketDistributionQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixEticketDistributionQueryRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.eticket.distribution.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixEticketDistributionQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参param
func (r *AlibabaDamaiMaitixEticketDistributionQueryRequest) SetParam(param *EticketQueryParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaDamaiMaitixEticketDistributionQueryRequest) GetParam() *EticketQueryParam {
    return r.param
}
