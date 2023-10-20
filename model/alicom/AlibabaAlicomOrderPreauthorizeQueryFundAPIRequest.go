package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest 资金流水查询 API请求
// alibaba.alicom.order.preauthorize.query.fund
//
// 预授权-资金流水查询
type AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest struct {
	model.Params
	// 入参
	_preAuthorizeModel *PreAuthorizeModel
}

// NewAlibabaAlicomOrderPreauthorizeQueryFundRequest 初始化AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest对象
func NewAlibabaAlicomOrderPreauthorizeQueryFundRequest() *AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest {
	return &AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.order.preauthorize.query.fund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPreAuthorizeModel is PreAuthorizeModel Setter
// 入参
func (r *AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest) SetPreAuthorizeModel(_preAuthorizeModel *PreAuthorizeModel) error {
	r._preAuthorizeModel = _preAuthorizeModel
	r.Set("pre_authorize_model", _preAuthorizeModel)
	return nil
}

// GetPreAuthorizeModel PreAuthorizeModel Getter
func (r AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest) GetPreAuthorizeModel() *PreAuthorizeModel {
	return r._preAuthorizeModel
}
