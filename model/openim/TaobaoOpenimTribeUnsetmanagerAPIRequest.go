package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimtribeunsetmanagerAPIRequest OPENIM群取消管理员 API请求
// taobao.openim.tribe.unsetmanager
//
// OPENIM群取消管理员
type TaobaoopenimtribeunsetmanagerAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tid int64
	// 用户信息
	_member *OpenImUser
}

// NewTaobaoopenimtribeunsetmanagerRequest 初始化TaobaoopenimtribeunsetmanagerAPIRequest对象
func NewTaobaoopenimtribeunsetmanagerRequest() *TaobaoopenimtribeunsetmanagerAPIRequest {
	return &TaobaoopenimtribeunsetmanagerAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenimtribeunsetmanagerAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.unsetmanager"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenimtribeunsetmanagerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenimtribeunsetmanagerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoopenimtribeunsetmanagerAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoopenimtribeunsetmanagerAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// SetTid is Tid Setter
// 群id
func (r *TaobaoopenimtribeunsetmanagerAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoopenimtribeunsetmanagerAPIRequest) GetTid() int64 {
	return r._tid
}

// SetMember is Member Setter
// 用户信息
func (r *TaobaoopenimtribeunsetmanagerAPIRequest) SetMember(_member *OpenImUser) error {
	r._member = _member
	r.Set("member", _member)
	return nil
}

// GetMember Member Getter
func (r TaobaoopenimtribeunsetmanagerAPIRequest) GetMember() *OpenImUser {
	return r._member
}
