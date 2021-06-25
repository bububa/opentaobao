package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发起流程 APIRequest
alibaba.mosflow.work.startprocess

业务发起流程审批
*/
type AlibabaMosflowWorkStartprocessRequest struct {
    model.Params

    // 参数二:额外业务参数,Map的JSON串
    variables   string 

    // 流程必传参数
    parameterEntity   *ParameterEntity 

}

func NewAlibabaMosflowWorkStartprocessRequest() *AlibabaMosflowWorkStartprocessRequest{
    return &AlibabaMosflowWorkStartprocessRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosflowWorkStartprocessRequest) GetApiMethodName() string {
    return "alibaba.mosflow.work.startprocess"
}

func (r AlibabaMosflowWorkStartprocessRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosflowWorkStartprocessRequest) SetVariables(variables string) error {
    r.variables = variables
    r.Set("variables", variables)
    return nil
}

func (r AlibabaMosflowWorkStartprocessRequest) GetVariables() string {
    return r.variables
}

func (r *AlibabaMosflowWorkStartprocessRequest) SetParameterEntity(parameterEntity *ParameterEntity) error {
    r.parameterEntity = parameterEntity
    r.Set("parameter_entity", parameterEntity)
    return nil
}

func (r AlibabaMosflowWorkStartprocessRequest) GetParameterEntity() *ParameterEntity {
    return r.parameterEntity
}

