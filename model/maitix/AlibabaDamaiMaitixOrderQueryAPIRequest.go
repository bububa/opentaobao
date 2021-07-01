package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦-查询分销单 API请求
alibaba.damai.maitix.order.query

查询分销单
*/
type AlibabaDamaiMaitixOrderQueryAPIRequest struct {
    model.Params
    // 分销单查询入参
    _param   *MoaOrderQueryParam
}

// 初始化AlibabaDamaiMaitixOrderQueryAPIRequest对象
func NewAlibabaDamaiMaitixOrderQueryRequest() *AlibabaDamaiMaitixOrderQueryAPIRequest{
    return &AlibabaDamaiMaitixOrderQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOrderQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.order.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOrderQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 分销单查询入参
func (r *AlibabaDamaiMaitixOrderQueryAPIRequest) SetParam(_param *MoaOrderQueryParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaDamaiMaitixOrderQueryAPIRequest) GetParam() *MoaOrderQueryParam {
    return r._param
}
