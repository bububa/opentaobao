package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦-新分销下单 API请求
alibaba.damai.maitix.order.distribution.create

createDistributionOrder
*/
type AlibabaDamaiMaitixOrderDistributionCreateRequest struct {
    model.Params
    // 下单参数param
    _param   *MoaOrderParam
}

// 初始化AlibabaDamaiMaitixOrderDistributionCreateRequest对象
func NewAlibabaDamaiMaitixOrderDistributionCreateRequest() *AlibabaDamaiMaitixOrderDistributionCreateRequest{
    return &AlibabaDamaiMaitixOrderDistributionCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOrderDistributionCreateRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.order.distribution.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOrderDistributionCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 下单参数param
func (r *AlibabaDamaiMaitixOrderDistributionCreateRequest) SetParam(_param *MoaOrderParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaDamaiMaitixOrderDistributionCreateRequest) GetParam() *MoaOrderParam {
    return r._param
}
