package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AccountAliyuncsComGetPubKey2013_07_01APIRequest 获取用户公钥 API请求
// account.aliyuncs.com.GetPubKey.2013-07-01
//
// 根据用户的appkey查询用户的pubkey
type AccountAliyuncsComGetPubKey2013_07_01APIRequest struct {
	model.Params
	// appkey
	_ownerAppkey string
}

// NewAccountAliyuncsComGetPubKey2013_07_01Request 初始化AccountAliyuncsComGetPubKey2013_07_01APIRequest对象
func NewAccountAliyuncsComGetPubKey2013_07_01Request() *AccountAliyuncsComGetPubKey2013_07_01APIRequest {
	return &AccountAliyuncsComGetPubKey2013_07_01APIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AccountAliyuncsComGetPubKey2013_07_01APIRequest) GetApiMethodName() string {
	return "account.aliyuncs.com.GetPubKey.2013-07-01"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AccountAliyuncsComGetPubKey2013_07_01APIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OwnerAppkey Setter
// appkey
func (r *AccountAliyuncsComGetPubKey2013_07_01APIRequest) SetOwnerAppkey(_ownerAppkey string) error {
	r._ownerAppkey = _ownerAppkey
	r.Set("OwnerAppkey", _ownerAppkey)
	return nil
}

// Get OwnerAppkey Getter
func (r AccountAliyuncsComGetPubKey2013_07_01APIRequest) GetOwnerAppkey() string {
	return r._ownerAppkey
}
