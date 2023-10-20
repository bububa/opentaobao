package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalicomorderpreauthorizequeryfundAPIRequest 资金流水查询 API请求
// alibaba.alicom.order.preauthorize.query.fund
//
// 预授权-资金流水查询
type AlibabaalicomorderpreauthorizequeryfundAPIRequest struct {
	model.Params
	// 入参
	_preAuthorizeModel *PreAuthorizeModel
}

// NewAlibabaalicomorderpreauthorizequeryfundRequest 初始化AlibabaalicomorderpreauthorizequeryfundAPIRequest对象
func NewAlibabaalicomorderpreauthorizequeryfundRequest() *AlibabaalicomorderpreauthorizequeryfundAPIRequest {
	return &AlibabaalicomorderpreauthorizequeryfundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalicomorderpreauthorizequeryfundAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.order.preauthorize.query.fund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalicomorderpreauthorizequeryfundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalicomorderpreauthorizequeryfundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPreAuthorizeModel is PreAuthorizeModel Setter
// 入参
func (r *AlibabaalicomorderpreauthorizequeryfundAPIRequest) SetPreAuthorizeModel(_preAuthorizeModel *PreAuthorizeModel) error {
	r._preAuthorizeModel = _preAuthorizeModel
	r.Set("pre_authorize_model", _preAuthorizeModel)
	return nil
}

// GetPreAuthorizeModel PreAuthorizeModel Getter
func (r AlibabaalicomorderpreauthorizequeryfundAPIRequest) GetPreAuthorizeModel() *PreAuthorizeModel {
	return r._preAuthorizeModel
}
