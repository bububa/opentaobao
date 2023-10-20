package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimtribeinviteAPIRequest OPENIM群邀请加入 API请求
// taobao.openim.tribe.invite
//
// OPENIM群邀请加入接口
type TaobaoopenimtribeinviteAPIRequest struct {
	model.Params
	// 用户信息
	_members []OpenImUser
	// 群id
	_tribeId int64
	// 用户信息
	_user *OpenImUser
}

// NewTaobaoopenimtribeinviteRequest 初始化TaobaoopenimtribeinviteAPIRequest对象
func NewTaobaoopenimtribeinviteRequest() *TaobaoopenimtribeinviteAPIRequest {
	return &TaobaoopenimtribeinviteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenimtribeinviteAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.invite"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenimtribeinviteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenimtribeinviteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMembers is Members Setter
// 用户信息
func (r *TaobaoopenimtribeinviteAPIRequest) SetMembers(_members []OpenImUser) error {
	r._members = _members
	r.Set("members", _members)
	return nil
}

// GetMembers Members Getter
func (r TaobaoopenimtribeinviteAPIRequest) GetMembers() []OpenImUser {
	return r._members
}

// SetTribeId is TribeId Setter
// 群id
func (r *TaobaoopenimtribeinviteAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// GetTribeId TribeId Getter
func (r TaobaoopenimtribeinviteAPIRequest) GetTribeId() int64 {
	return r._tribeId
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoopenimtribeinviteAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoopenimtribeinviteAPIRequest) GetUser() *OpenImUser {
	return r._user
}
