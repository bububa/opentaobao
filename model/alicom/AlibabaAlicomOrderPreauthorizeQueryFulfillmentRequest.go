package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
履约结果查询 APIRequest
alibaba.alicom.order.preauthorize.query.fulfillment

预授权-履约结果查询
*/
type AlibabaAlicomOrderPreauthorizeQueryFulfillmentRequest struct {
    model.Params

    // 入参
    preAuthorizeModel   *PreAuthorizeModel 

}

func NewAlibabaAlicomOrderPreauthorizeQueryFulfillmentRequest() *AlibabaAlicomOrderPreauthorizeQueryFulfillmentRequest{
    return &AlibabaAlicomOrderPreauthorizeQueryFulfillmentRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlicomOrderPreauthorizeQueryFulfillmentRequest) GetApiMethodName() string {
    return "alibaba.alicom.order.preauthorize.query.fulfillment"
}

func (r AlibabaAlicomOrderPreauthorizeQueryFulfillmentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlicomOrderPreauthorizeQueryFulfillmentRequest) SetPreAuthorizeModel(preAuthorizeModel *PreAuthorizeModel) error {
    r.preAuthorizeModel = preAuthorizeModel
    r.Set("pre_authorize_model", preAuthorizeModel)
    return nil
}

func (r AlibabaAlicomOrderPreauthorizeQueryFulfillmentRequest) GetPreAuthorizeModel() *PreAuthorizeModel {
    return r.preAuthorizeModel
}

