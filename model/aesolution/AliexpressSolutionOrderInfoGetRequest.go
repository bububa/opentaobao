package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
get order detail info API请求
aliexpress.solution.order.info.get

get order detail info
*/
type AliexpressSolutionOrderInfoGetRequest struct {
    model.Params
    // param
    _param1   *OrderDetailQuery
}

// 初始化AliexpressSolutionOrderInfoGetRequest对象
func NewAliexpressSolutionOrderInfoGetRequest() *AliexpressSolutionOrderInfoGetRequest{
    return &AliexpressSolutionOrderInfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionOrderInfoGetRequest) GetApiMethodName() string {
    return "aliexpress.solution.order.info.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionOrderInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param1 Setter
// param
func (r *AliexpressSolutionOrderInfoGetRequest) SetParam1(_param1 *OrderDetailQuery) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AliexpressSolutionOrderInfoGetRequest) GetParam1() *OrderDetailQuery {
    return r._param1
}
