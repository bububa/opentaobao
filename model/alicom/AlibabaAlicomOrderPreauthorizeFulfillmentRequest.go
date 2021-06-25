package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
履约结果 APIRequest
alibaba.alicom.order.preauthorize.fulfillment

预授权-履约结果
*/
type AlibabaAlicomOrderPreauthorizeFulfillmentRequest struct {
    model.Params

    // 入参
    preAuthorizeModel   *PreAuthorizeModel 

}

func NewAlibabaAlicomOrderPreauthorizeFulfillmentRequest() *AlibabaAlicomOrderPreauthorizeFulfillmentRequest{
    return &AlibabaAlicomOrderPreauthorizeFulfillmentRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlicomOrderPreauthorizeFulfillmentRequest) GetApiMethodName() string {
    return "alibaba.alicom.order.preauthorize.fulfillment"
}

func (r AlibabaAlicomOrderPreauthorizeFulfillmentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlicomOrderPreauthorizeFulfillmentRequest) SetPreAuthorizeModel(preAuthorizeModel *PreAuthorizeModel) error {
    r.preAuthorizeModel = preAuthorizeModel
    r.Set("pre_authorize_model", preAuthorizeModel)
    return nil
}

func (r AlibabaAlicomOrderPreauthorizeFulfillmentRequest) GetPreAuthorizeModel() *PreAuthorizeModel {
    return r.preAuthorizeModel
}

