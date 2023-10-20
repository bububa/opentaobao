package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimtribesetmanagerAPIRequest OPENIM群设置管理员 API请求
// taobao.openim.tribe.setmanager
//
// OPENIM群设置管理员
type TaobaoopenimtribesetmanagerAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tid int64
	// 用户信息
	_member *OpenImUser
}

// NewTaobaoopenimtribesetmanagerRequest 初始化TaobaoopenimtribesetmanagerAPIRequest对象
func NewTaobaoopenimtribesetmanagerRequest() *TaobaoopenimtribesetmanagerAPIRequest {
	return &TaobaoopenimtribesetmanagerAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenimtribesetmanagerAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.setmanager"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenimtribesetmanagerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenimtribesetmanagerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoopenimtribesetmanagerAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoopenimtribesetmanagerAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// SetTid is Tid Setter
// 群id
func (r *TaobaoopenimtribesetmanagerAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoopenimtribesetmanagerAPIRequest) GetTid() int64 {
	return r._tid
}

// SetMember is Member Setter
// 用户信息
func (r *TaobaoopenimtribesetmanagerAPIRequest) SetMember(_member *OpenImUser) error {
	r._member = _member
	r.Set("member", _member)
	return nil
}

// GetMember Member Getter
func (r TaobaoopenimtribesetmanagerAPIRequest) GetMember() *OpenImUser {
	return r._member
}
