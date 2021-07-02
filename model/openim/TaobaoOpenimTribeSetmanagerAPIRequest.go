package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeSetmanagerAPIRequest OPENIM群设置管理员 API请求
// taobao.openim.tribe.setmanager
//
// OPENIM群设置管理员
type TaobaoOpenimTribeSetmanagerAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tid int64
	// 用户信息
	_member *OpenImUser
}

// NewTaobaoOpenimTribeSetmanagerRequest 初始化TaobaoOpenimTribeSetmanagerAPIRequest对象
func NewTaobaoOpenimTribeSetmanagerRequest() *TaobaoOpenimTribeSetmanagerAPIRequest {
	return &TaobaoOpenimTribeSetmanagerAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeSetmanagerAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.setmanager"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeSetmanagerAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is User Setter
// 用户信息
func (r *TaobaoOpenimTribeSetmanagerAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// Get User Getter
func (r TaobaoOpenimTribeSetmanagerAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// Set is Tid Setter
// 群id
func (r *TaobaoOpenimTribeSetmanagerAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// Get Tid Getter
func (r TaobaoOpenimTribeSetmanagerAPIRequest) GetTid() int64 {
	return r._tid
}

// Set is Member Setter
// 用户信息
func (r *TaobaoOpenimTribeSetmanagerAPIRequest) SetMember(_member *OpenImUser) error {
	r._member = _member
	r.Set("member", _member)
	return nil
}

// Get Member Getter
func (r TaobaoOpenimTribeSetmanagerAPIRequest) GetMember() *OpenImUser {
	return r._member
}
