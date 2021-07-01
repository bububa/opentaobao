package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlicomOrderPreauthorizeCreateAPIRequest
业务办理结果 API请求
alibaba.alicom.order.preauthorize.create

授授权:签约结果通知 */
type AlibabaAlicomOrderPreauthorizeCreateAPIRequest struct {
	model.Params
	// 入参
	_preAuthorizeModel *PreAuthorizeModel
}

// NewAlibabaAlicomOrderPreauthorizeCreateRequest 初始化AlibabaAlicomOrderPreauthorizeCreateAPIRequest对象
func NewAlibabaAlicomOrderPreauthorizeCreateRequest() *AlibabaAlicomOrderPreauthorizeCreateAPIRequest {
	return &AlibabaAlicomOrderPreauthorizeCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlicomOrderPreauthorizeCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.order.preauthorize.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlicomOrderPreauthorizeCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PreAuthorizeModel Setter
// 入参
func (r *AlibabaAlicomOrderPreauthorizeCreateAPIRequest) SetPreAuthorizeModel(_preAuthorizeModel *PreAuthorizeModel) error {
	r._preAuthorizeModel = _preAuthorizeModel
	r.Set("pre_authorize_model", _preAuthorizeModel)
	return nil
}

// Get PreAuthorizeModel Getter
func (r AlibabaAlicomOrderPreauthorizeCreateAPIRequest) GetPreAuthorizeModel() *PreAuthorizeModel {
	return r._preAuthorizeModel
}
