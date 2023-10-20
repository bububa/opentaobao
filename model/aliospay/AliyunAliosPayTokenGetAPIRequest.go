package aliospay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunAliosPayTokenGetAPIRequest 获取支付token API请求
// aliyun.alios.pay.token.get
//
// 商户用来获取支付的授权token
type AliyunAliosPayTokenGetAPIRequest struct {
	model.Params
	// 请求参数
	_getTokenRequest *GetTokenRequest
}

// NewAliyunAliosPayTokenGetRequest 初始化AliyunAliosPayTokenGetAPIRequest对象
func NewAliyunAliosPayTokenGetRequest() *AliyunAliosPayTokenGetAPIRequest {
	return &AliyunAliosPayTokenGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunAliosPayTokenGetAPIRequest) GetApiMethodName() string {
	return "aliyun.alios.pay.token.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunAliosPayTokenGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunAliosPayTokenGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGetTokenRequest is GetTokenRequest Setter
// 请求参数
func (r *AliyunAliosPayTokenGetAPIRequest) SetGetTokenRequest(_getTokenRequest *GetTokenRequest) error {
	r._getTokenRequest = _getTokenRequest
	r.Set("get_token_request", _getTokenRequest)
	return nil
}

// GetGetTokenRequest GetTokenRequest Getter
func (r AliyunAliosPayTokenGetAPIRequest) GetGetTokenRequest() *GetTokenRequest {
	return r._getTokenRequest
}
