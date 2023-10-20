package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AccountaliyuncscomgetPubKey20130701APIRequest 获取用户公钥 API请求
// account.aliyuncs.com.GetPubKey.2013-07-01
//
// 根据用户的appkey查询用户的pubkey
type AccountaliyuncscomgetPubKey20130701APIRequest struct {
	model.Params
	// appkey
	_ownerAppkey string
}

// NewAccountaliyuncscomgetPubKey20130701Request 初始化AccountaliyuncscomgetPubKey20130701APIRequest对象
func NewAccountaliyuncscomgetPubKey20130701Request() *AccountaliyuncscomgetPubKey20130701APIRequest {
	return &AccountaliyuncscomgetPubKey20130701APIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AccountaliyuncscomgetPubKey20130701APIRequest) GetApiMethodName() string {
	return "account.aliyuncs.com.GetPubKey.2013-07-01"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AccountaliyuncscomgetPubKey20130701APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AccountaliyuncscomgetPubKey20130701APIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOwnerAppkey is OwnerAppkey Setter
// appkey
func (r *AccountaliyuncscomgetPubKey20130701APIRequest) SetOwnerAppkey(_ownerAppkey string) error {
	r._ownerAppkey = _ownerAppkey
	r.Set("OwnerAppkey", _ownerAppkey)
	return nil
}

// GetOwnerAppkey OwnerAppkey Getter
func (r AccountaliyuncscomgetPubKey20130701APIRequest) GetOwnerAppkey() string {
	return r._ownerAppkey
}
