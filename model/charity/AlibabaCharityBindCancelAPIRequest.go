package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityBindCancelAPIRequest 取消用户绑定 API请求
// alibaba.charity.bind.cancel
//
// 取消用户绑定
type AlibabaCharityBindCancelAPIRequest struct {
	model.Params
	// 解绑用户ID
	_userKey string
}

// NewAlibabaCharityBindCancelRequest 初始化AlibabaCharityBindCancelAPIRequest对象
func NewAlibabaCharityBindCancelRequest() *AlibabaCharityBindCancelAPIRequest {
	return &AlibabaCharityBindCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCharityBindCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.charity.bind.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCharityBindCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUserKey is UserKey Setter
// 解绑用户ID
func (r *AlibabaCharityBindCancelAPIRequest) SetUserKey(_userKey string) error {
	r._userKey = _userKey
	r.Set("user_key", _userKey)
	return nil
}

// GetUserKey UserKey Getter
func (r AlibabaCharityBindCancelAPIRequest) GetUserKey() string {
	return r._userKey
}
