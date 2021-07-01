package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发起流程 API请求
alibaba.mosflow.work.startprocess

业务发起流程审批
*/
type AlibabaMosflowWorkStartprocessAPIRequest struct {
    model.Params
    // 参数二:额外业务参数,Map的JSON串
    _variables   string
    // 流程必传参数
    _parameterEntity   *ParameterEntity
}

// 初始化AlibabaMosflowWorkStartprocessAPIRequest对象
func NewAlibabaMosflowWorkStartprocessRequest() *AlibabaMosflowWorkStartprocessAPIRequest{
    return &AlibabaMosflowWorkStartprocessAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosflowWorkStartprocessAPIRequest) GetApiMethodName() string {
    return "alibaba.mosflow.work.startprocess"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosflowWorkStartprocessAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Variables Setter
// 参数二:额外业务参数,Map的JSON串
func (r *AlibabaMosflowWorkStartprocessAPIRequest) SetVariables(_variables string) error {
    r._variables = _variables
    r.Set("variables", _variables)
    return nil
}

// Variables Getter
func (r AlibabaMosflowWorkStartprocessAPIRequest) GetVariables() string {
    return r._variables
}
// ParameterEntity Setter
// 流程必传参数
func (r *AlibabaMosflowWorkStartprocessAPIRequest) SetParameterEntity(_parameterEntity *ParameterEntity) error {
    r._parameterEntity = _parameterEntity
    r.Set("parameter_entity", _parameterEntity)
    return nil
}

// ParameterEntity Getter
func (r AlibabaMosflowWorkStartprocessAPIRequest) GetParameterEntity() *ParameterEntity {
    return r._parameterEntity
}
