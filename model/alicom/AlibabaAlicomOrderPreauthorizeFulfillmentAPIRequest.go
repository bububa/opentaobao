package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlicomOrderPreauthorizeFulfillmentAPIRequest 履约结果 API请求
// alibaba.alicom.order.preauthorize.fulfillment
//
// 预授权-履约结果
type AlibabaAlicomOrderPreauthorizeFulfillmentAPIRequest struct {
	model.Params
	// 入参
	_preAuthorizeModel *PreAuthorizeModel
}

// NewAlibabaAlicomOrderPreauthorizeFulfillmentRequest 初始化AlibabaAlicomOrderPreauthorizeFulfillmentAPIRequest对象
func NewAlibabaAlicomOrderPreauthorizeFulfillmentRequest() *AlibabaAlicomOrderPreauthorizeFulfillmentAPIRequest {
	return &AlibabaAlicomOrderPreauthorizeFulfillmentAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlicomOrderPreauthorizeFulfillmentAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.order.preauthorize.fulfillment"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlicomOrderPreauthorizeFulfillmentAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PreAuthorizeModel Setter
// 入参
func (r *AlibabaAlicomOrderPreauthorizeFulfillmentAPIRequest) SetPreAuthorizeModel(_preAuthorizeModel *PreAuthorizeModel) error {
	r._preAuthorizeModel = _preAuthorizeModel
	r.Set("pre_authorize_model", _preAuthorizeModel)
	return nil
}

// Get PreAuthorizeModel Getter
func (r AlibabaAlicomOrderPreauthorizeFulfillmentAPIRequest) GetPreAuthorizeModel() *PreAuthorizeModel {
	return r._preAuthorizeModel
}
