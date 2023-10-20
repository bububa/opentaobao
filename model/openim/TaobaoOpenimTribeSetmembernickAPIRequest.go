package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimtribesetmembernickAPIRequest 设置群成员昵称 API请求
// taobao.openim.tribe.setmembernick
//
// 设置群成员昵称，存在如下两种场景
// 1 群主或管理员设置群成员昵称，该操作有权限控制。只针对普通群的群主和管理员开发此功能；讨论组群主不支持此设置操作
// 2 群成员设置自己的昵称，该功能对群所有成员开放
type TaobaoopenimtribesetmembernickAPIRequest struct {
	model.Params
	// 设置的昵称
	_nick string
	// 发起设置昵称的操作者，如果是设置其他成员的昵称，只有普通组的群主和管理员有权限
	_user *User
	// 群id
	_tribeId int64
	// 被设置昵称的群成员
	_member *User
}

// NewTaobaoopenimtribesetmembernickRequest 初始化TaobaoopenimtribesetmembernickAPIRequest对象
func NewTaobaoopenimtribesetmembernickRequest() *TaobaoopenimtribesetmembernickAPIRequest {
	return &TaobaoopenimtribesetmembernickAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenimtribesetmembernickAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.setmembernick"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenimtribesetmembernickAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenimtribesetmembernickAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 设置的昵称
func (r *TaobaoopenimtribesetmembernickAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoopenimtribesetmembernickAPIRequest) GetNick() string {
	return r._nick
}

// SetUser is User Setter
// 发起设置昵称的操作者，如果是设置其他成员的昵称，只有普通组的群主和管理员有权限
func (r *TaobaoopenimtribesetmembernickAPIRequest) SetUser(_user *User) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoopenimtribesetmembernickAPIRequest) GetUser() *User {
	return r._user
}

// SetTribeId is TribeId Setter
// 群id
func (r *TaobaoopenimtribesetmembernickAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// GetTribeId TribeId Getter
func (r TaobaoopenimtribesetmembernickAPIRequest) GetTribeId() int64 {
	return r._tribeId
}

// SetMember is Member Setter
// 被设置昵称的群成员
func (r *TaobaoopenimtribesetmembernickAPIRequest) SetMember(_member *User) error {
	r._member = _member
	r.Set("member", _member)
	return nil
}

// GetMember Member Getter
func (r TaobaoopenimtribesetmembernickAPIRequest) GetMember() *User {
	return r._member
}
