package charity

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCharityBindCancelAPIRequest) Reset() {
	r._userKey = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCharityBindCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.charity.bind.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCharityBindCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCharityBindCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaCharityBindCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCharityBindCancelRequest()
	},
}

// GetAlibabaCharityBindCancelRequest 从 sync.Pool 获取 AlibabaCharityBindCancelAPIRequest
func GetAlibabaCharityBindCancelAPIRequest() *AlibabaCharityBindCancelAPIRequest {
	return poolAlibabaCharityBindCancelAPIRequest.Get().(*AlibabaCharityBindCancelAPIRequest)
}

// ReleaseAlibabaCharityBindCancelAPIRequest 将 AlibabaCharityBindCancelAPIRequest 放入 sync.Pool
func ReleaseAlibabaCharityBindCancelAPIRequest(v *AlibabaCharityBindCancelAPIRequest) {
	v.Reset()
	poolAlibabaCharityBindCancelAPIRequest.Put(v)
}
