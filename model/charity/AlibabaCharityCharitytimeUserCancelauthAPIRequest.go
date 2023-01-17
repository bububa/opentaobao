package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityCharitytimeUserCancelauthAPIRequest 取消用户授权 API请求
// alibaba.charity.charitytime.user.cancelauth
//
// 取消用户授权
type AlibabaCharityCharitytimeUserCancelauthAPIRequest struct {
	model.Params
	// 参数对象
	_cancelAuthHsfRequest *CancelAuthHsfRequest
}

// NewAlibabaCharityCharitytimeUserCancelauthRequest 初始化AlibabaCharityCharitytimeUserCancelauthAPIRequest对象
func NewAlibabaCharityCharitytimeUserCancelauthRequest() *AlibabaCharityCharitytimeUserCancelauthAPIRequest {
	return &AlibabaCharityCharitytimeUserCancelauthAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCharityCharitytimeUserCancelauthAPIRequest) GetApiMethodName() string {
	return "alibaba.charity.charitytime.user.cancelauth"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCharityCharitytimeUserCancelauthAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCharityCharitytimeUserCancelauthAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCancelAuthHsfRequest is CancelAuthHsfRequest Setter
// 参数对象
func (r *AlibabaCharityCharitytimeUserCancelauthAPIRequest) SetCancelAuthHsfRequest(_cancelAuthHsfRequest *CancelAuthHsfRequest) error {
	r._cancelAuthHsfRequest = _cancelAuthHsfRequest
	r.Set("cancel_auth_hsf_request", _cancelAuthHsfRequest)
	return nil
}

// GetCancelAuthHsfRequest CancelAuthHsfRequest Getter
func (r AlibabaCharityCharitytimeUserCancelauthAPIRequest) GetCancelAuthHsfRequest() *CancelAuthHsfRequest {
	return r._cancelAuthHsfRequest
}
