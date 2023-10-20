package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlicomOrderPreauthorizeCreateAPIRequest 业务办理结果 API请求
// alibaba.alicom.order.preauthorize.create
//
// 授授权:签约结果通知
type AlibabaAlicomOrderPreauthorizeCreateAPIRequest struct {
	model.Params
	// 入参
	_preAuthorizeModel *PreAuthorizeModel
}

// NewAlibabaAlicomOrderPreauthorizeCreateRequest 初始化AlibabaAlicomOrderPreauthorizeCreateAPIRequest对象
func NewAlibabaAlicomOrderPreauthorizeCreateRequest() *AlibabaAlicomOrderPreauthorizeCreateAPIRequest {
	return &AlibabaAlicomOrderPreauthorizeCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlicomOrderPreauthorizeCreateAPIRequest) Reset() {
	r._preAuthorizeModel = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlicomOrderPreauthorizeCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.order.preauthorize.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlicomOrderPreauthorizeCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlicomOrderPreauthorizeCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPreAuthorizeModel is PreAuthorizeModel Setter
// 入参
func (r *AlibabaAlicomOrderPreauthorizeCreateAPIRequest) SetPreAuthorizeModel(_preAuthorizeModel *PreAuthorizeModel) error {
	r._preAuthorizeModel = _preAuthorizeModel
	r.Set("pre_authorize_model", _preAuthorizeModel)
	return nil
}

// GetPreAuthorizeModel PreAuthorizeModel Getter
func (r AlibabaAlicomOrderPreauthorizeCreateAPIRequest) GetPreAuthorizeModel() *PreAuthorizeModel {
	return r._preAuthorizeModel
}

var poolAlibabaAlicomOrderPreauthorizeCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlicomOrderPreauthorizeCreateRequest()
	},
}

// GetAlibabaAlicomOrderPreauthorizeCreateRequest 从 sync.Pool 获取 AlibabaAlicomOrderPreauthorizeCreateAPIRequest
func GetAlibabaAlicomOrderPreauthorizeCreateAPIRequest() *AlibabaAlicomOrderPreauthorizeCreateAPIRequest {
	return poolAlibabaAlicomOrderPreauthorizeCreateAPIRequest.Get().(*AlibabaAlicomOrderPreauthorizeCreateAPIRequest)
}

// ReleaseAlibabaAlicomOrderPreauthorizeCreateAPIRequest 将 AlibabaAlicomOrderPreauthorizeCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlicomOrderPreauthorizeCreateAPIRequest(v *AlibabaAlicomOrderPreauthorizeCreateAPIRequest) {
	v.Reset()
	poolAlibabaAlicomOrderPreauthorizeCreateAPIRequest.Put(v)
}
