package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
业务办理结果 APIRequest
alibaba.alicom.order.preauthorize.create

授授权:签约结果通知
*/
type AlibabaAlicomOrderPreauthorizeCreateRequest struct {
    model.Params

    // 入参
    preAuthorizeModel   *PreAuthorizeModel 

}

func NewAlibabaAlicomOrderPreauthorizeCreateRequest() *AlibabaAlicomOrderPreauthorizeCreateRequest{
    return &AlibabaAlicomOrderPreauthorizeCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlicomOrderPreauthorizeCreateRequest) GetApiMethodName() string {
    return "alibaba.alicom.order.preauthorize.create"
}

func (r AlibabaAlicomOrderPreauthorizeCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlicomOrderPreauthorizeCreateRequest) SetPreAuthorizeModel(preAuthorizeModel *PreAuthorizeModel) error {
    r.preAuthorizeModel = preAuthorizeModel
    r.Set("pre_authorize_model", preAuthorizeModel)
    return nil
}

func (r AlibabaAlicomOrderPreauthorizeCreateRequest) GetPreAuthorizeModel() *PreAuthorizeModel {
    return r.preAuthorizeModel
}

