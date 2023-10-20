package user

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AccountAliyuncsComGetPubKey20130701APIRequest 获取用户公钥 API请求
// account.aliyuncs.com.GetPubKey.2013-07-01
//
// 根据用户的appkey查询用户的pubkey
type AccountAliyuncsComGetPubKey20130701APIRequest struct {
	model.Params
	// appkey
	_ownerAppkey string
}

// NewAccountAliyuncsComGetPubKey20130701Request 初始化AccountAliyuncsComGetPubKey20130701APIRequest对象
func NewAccountAliyuncsComGetPubKey20130701Request() *AccountAliyuncsComGetPubKey20130701APIRequest {
	return &AccountAliyuncsComGetPubKey20130701APIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AccountAliyuncsComGetPubKey20130701APIRequest) Reset() {
	r._ownerAppkey = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AccountAliyuncsComGetPubKey20130701APIRequest) GetApiMethodName() string {
	return "account.aliyuncs.com.GetPubKey.2013-07-01"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AccountAliyuncsComGetPubKey20130701APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AccountAliyuncsComGetPubKey20130701APIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOwnerAppkey is OwnerAppkey Setter
// appkey
func (r *AccountAliyuncsComGetPubKey20130701APIRequest) SetOwnerAppkey(_ownerAppkey string) error {
	r._ownerAppkey = _ownerAppkey
	r.Set("OwnerAppkey", _ownerAppkey)
	return nil
}

// GetOwnerAppkey OwnerAppkey Getter
func (r AccountAliyuncsComGetPubKey20130701APIRequest) GetOwnerAppkey() string {
	return r._ownerAppkey
}

var poolAccountAliyuncsComGetPubKey20130701APIRequest = sync.Pool{
	New: func() any {
		return NewAccountAliyuncsComGetPubKey20130701Request()
	},
}

// GetAccountAliyuncsComGetPubKey20130701Request 从 sync.Pool 获取 AccountAliyuncsComGetPubKey20130701APIRequest
func GetAccountAliyuncsComGetPubKey20130701APIRequest() *AccountAliyuncsComGetPubKey20130701APIRequest {
	return poolAccountAliyuncsComGetPubKey20130701APIRequest.Get().(*AccountAliyuncsComGetPubKey20130701APIRequest)
}

// ReleaseAccountAliyuncsComGetPubKey20130701APIRequest 将 AccountAliyuncsComGetPubKey20130701APIRequest 放入 sync.Pool
func ReleaseAccountAliyuncsComGetPubKey20130701APIRequest(v *AccountAliyuncsComGetPubKey20130701APIRequest) {
	v.Reset()
	poolAccountAliyuncsComGetPubKey20130701APIRequest.Put(v)
}
