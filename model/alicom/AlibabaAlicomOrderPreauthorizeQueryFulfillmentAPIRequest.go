package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlicomOrderPreauthorizeQueryFulfillmentAPIRequest 履约结果查询 API请求
// alibaba.alicom.order.preauthorize.query.fulfillment
//
// 预授权-履约结果查询
type AlibabaAlicomOrderPreauthorizeQueryFulfillmentAPIRequest struct {
	model.Params
	// 入参
	_preAuthorizeModel *PreAuthorizeModel
}

// NewAlibabaAlicomOrderPreauthorizeQueryFulfillmentRequest 初始化AlibabaAlicomOrderPreauthorizeQueryFulfillmentAPIRequest对象
func NewAlibabaAlicomOrderPreauthorizeQueryFulfillmentRequest() *AlibabaAlicomOrderPreauthorizeQueryFulfillmentAPIRequest {
	return &AlibabaAlicomOrderPreauthorizeQueryFulfillmentAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlicomOrderPreauthorizeQueryFulfillmentAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.order.preauthorize.query.fulfillment"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlicomOrderPreauthorizeQueryFulfillmentAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPreAuthorizeModel is PreAuthorizeModel Setter
// 入参
func (r *AlibabaAlicomOrderPreauthorizeQueryFulfillmentAPIRequest) SetPreAuthorizeModel(_preAuthorizeModel *PreAuthorizeModel) error {
	r._preAuthorizeModel = _preAuthorizeModel
	r.Set("pre_authorize_model", _preAuthorizeModel)
	return nil
}

// GetPreAuthorizeModel PreAuthorizeModel Getter
func (r AlibabaAlicomOrderPreauthorizeQueryFulfillmentAPIRequest) GetPreAuthorizeModel() *PreAuthorizeModel {
	return r._preAuthorizeModel
}
