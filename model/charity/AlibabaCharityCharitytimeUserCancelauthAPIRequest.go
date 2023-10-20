package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacharitycharitytimeusercancelauthAPIRequest 取消用户授权 API请求
// alibaba.charity.charitytime.user.cancelauth
//
// 取消用户授权
type AlibabacharitycharitytimeusercancelauthAPIRequest struct {
	model.Params
	// 参数对象
	_cancelAuthHsfRequest *CancelAuthHsfRequest
}

// NewAlibabacharitycharitytimeusercancelauthRequest 初始化AlibabacharitycharitytimeusercancelauthAPIRequest对象
func NewAlibabacharitycharitytimeusercancelauthRequest() *AlibabacharitycharitytimeusercancelauthAPIRequest {
	return &AlibabacharitycharitytimeusercancelauthAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacharitycharitytimeusercancelauthAPIRequest) GetApiMethodName() string {
	return "alibaba.charity.charitytime.user.cancelauth"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacharitycharitytimeusercancelauthAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacharitycharitytimeusercancelauthAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCancelAuthHsfRequest is CancelAuthHsfRequest Setter
// 参数对象
func (r *AlibabacharitycharitytimeusercancelauthAPIRequest) SetCancelAuthHsfRequest(_cancelAuthHsfRequest *CancelAuthHsfRequest) error {
	r._cancelAuthHsfRequest = _cancelAuthHsfRequest
	r.Set("cancel_auth_hsf_request", _cancelAuthHsfRequest)
	return nil
}

// GetCancelAuthHsfRequest CancelAuthHsfRequest Getter
func (r AlibabacharitycharitytimeusercancelauthAPIRequest) GetCancelAuthHsfRequest() *CancelAuthHsfRequest {
	return r._cancelAuthHsfRequest
}
