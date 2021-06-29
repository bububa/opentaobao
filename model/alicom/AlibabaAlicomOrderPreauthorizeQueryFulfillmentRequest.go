package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
履约结果查询 API请求
alibaba.alicom.order.preauthorize.query.fulfillment

预授权-履约结果查询
*/
type AlibabaAlicomOrderPreauthorizeQueryFulfillmentRequest struct {
    model.Params
    // 入参
    preAuthorizeModel   *PreAuthorizeModel
}

// 初始化AlibabaAlicomOrderPreauthorizeQueryFulfillmentRequest对象
func NewAlibabaAlicomOrderPreauthorizeQueryFulfillmentRequest() *AlibabaAlicomOrderPreauthorizeQueryFulfillmentRequest{
    return &AlibabaAlicomOrderPreauthorizeQueryFulfillmentRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlicomOrderPreauthorizeQueryFulfillmentRequest) GetApiMethodName() string {
    return "alibaba.alicom.order.preauthorize.query.fulfillment"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlicomOrderPreauthorizeQueryFulfillmentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PreAuthorizeModel Setter
// 入参
func (r *AlibabaAlicomOrderPreauthorizeQueryFulfillmentRequest) SetPreAuthorizeModel(preAuthorizeModel *PreAuthorizeModel) error {
    r.preAuthorizeModel = preAuthorizeModel
    r.Set("pre_authorize_model", preAuthorizeModel)
    return nil
}

// PreAuthorizeModel Getter
func (r AlibabaAlicomOrderPreauthorizeQueryFulfillmentRequest) GetPreAuthorizeModel() *PreAuthorizeModel {
    return r.preAuthorizeModel
}
