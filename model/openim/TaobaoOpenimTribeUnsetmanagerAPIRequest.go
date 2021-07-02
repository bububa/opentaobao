package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeUnsetmanagerAPIRequest OPENIM群取消管理员 API请求
// taobao.openim.tribe.unsetmanager
//
// OPENIM群取消管理员
type TaobaoOpenimTribeUnsetmanagerAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tid int64
	// 用户信息
	_member *OpenImUser
}

// NewTaobaoOpenimTribeUnsetmanagerRequest 初始化TaobaoOpenimTribeUnsetmanagerAPIRequest对象
func NewTaobaoOpenimTribeUnsetmanagerRequest() *TaobaoOpenimTribeUnsetmanagerAPIRequest {
	return &TaobaoOpenimTribeUnsetmanagerAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeUnsetmanagerAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.unsetmanager"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeUnsetmanagerAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is User Setter
// 用户信息
func (r *TaobaoOpenimTribeUnsetmanagerAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// Get User Getter
func (r TaobaoOpenimTribeUnsetmanagerAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// Set is Tid Setter
// 群id
func (r *TaobaoOpenimTribeUnsetmanagerAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// Get Tid Getter
func (r TaobaoOpenimTribeUnsetmanagerAPIRequest) GetTid() int64 {
	return r._tid
}

// Set is Member Setter
// 用户信息
func (r *TaobaoOpenimTribeUnsetmanagerAPIRequest) SetMember(_member *OpenImUser) error {
	r._member = _member
	r.Set("member", _member)
	return nil
}

// Get Member Getter
func (r TaobaoOpenimTribeUnsetmanagerAPIRequest) GetMember() *OpenImUser {
	return r._member
}
