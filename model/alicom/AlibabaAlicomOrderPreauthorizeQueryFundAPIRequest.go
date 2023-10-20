package alicom

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest) Reset() {
	r._preAuthorizeModel = nil
	r.Params.ToZero()
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

var poolAlibabaAlicomOrderPreauthorizeQueryFundAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlicomOrderPreauthorizeQueryFundRequest()
	},
}

// GetAlibabaAlicomOrderPreauthorizeQueryFundRequest 从 sync.Pool 获取 AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest
func GetAlibabaAlicomOrderPreauthorizeQueryFundAPIRequest() *AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest {
	return poolAlibabaAlicomOrderPreauthorizeQueryFundAPIRequest.Get().(*AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest)
}

// ReleaseAlibabaAlicomOrderPreauthorizeQueryFundAPIRequest 将 AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlicomOrderPreauthorizeQueryFundAPIRequest(v *AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest) {
	v.Reset()
	poolAlibabaAlicomOrderPreauthorizeQueryFundAPIRequest.Put(v)
}
