package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacharitybindcancelAPIRequest 取消用户绑定 API请求
// alibaba.charity.bind.cancel
//
// 取消用户绑定
type AlibabacharitybindcancelAPIRequest struct {
	model.Params
	// 解绑用户ID
	_userKey string
}

// NewAlibabacharitybindcancelRequest 初始化AlibabacharitybindcancelAPIRequest对象
func NewAlibabacharitybindcancelRequest() *AlibabacharitybindcancelAPIRequest {
	return &AlibabacharitybindcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacharitybindcancelAPIRequest) GetApiMethodName() string {
	return "alibaba.charity.bind.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacharitybindcancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacharitybindcancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserKey is UserKey Setter
// 解绑用户ID
func (r *AlibabacharitybindcancelAPIRequest) SetUserKey(_userKey string) error {
	r._userKey = _userKey
	r.Set("user_key", _userKey)
	return nil
}

// GetUserKey UserKey Getter
func (r AlibabacharitybindcancelAPIRequest) GetUserKey() string {
	return r._userKey
}
