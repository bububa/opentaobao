package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
履约结果 API请求
alibaba.alicom.order.preauthorize.fulfillment

预授权-履约结果
*/
type AlibabaAlicomOrderPreauthorizeFulfillmentRequest struct {
    model.Params
    // 入参
    preAuthorizeModel   *PreAuthorizeModel
}

// 初始化AlibabaAlicomOrderPreauthorizeFulfillmentRequest对象
func NewAlibabaAlicomOrderPreauthorizeFulfillmentRequest() *AlibabaAlicomOrderPreauthorizeFulfillmentRequest{
    return &AlibabaAlicomOrderPreauthorizeFulfillmentRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlicomOrderPreauthorizeFulfillmentRequest) GetApiMethodName() string {
    return "alibaba.alicom.order.preauthorize.fulfillment"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlicomOrderPreauthorizeFulfillmentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PreAuthorizeModel Setter
// 入参
func (r *AlibabaAlicomOrderPreauthorizeFulfillmentRequest) SetPreAuthorizeModel(preAuthorizeModel *PreAuthorizeModel) error {
    r.preAuthorizeModel = preAuthorizeModel
    r.Set("pre_authorize_model", preAuthorizeModel)
    return nil
}

// PreAuthorizeModel Getter
func (r AlibabaAlicomOrderPreauthorizeFulfillmentRequest) GetPreAuthorizeModel() *PreAuthorizeModel {
    return r.preAuthorizeModel
}
