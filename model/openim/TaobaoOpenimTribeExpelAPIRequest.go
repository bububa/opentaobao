package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimtribeexpelAPIRequest OPENIM群踢出成员 API请求
// taobao.openim.tribe.expel
//
// OPENIM群踢出成员
type TaobaoopenimtribeexpelAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tribeId int64
	// 用户信息
	_member *OpenImUser
}

// NewTaobaoopenimtribeexpelRequest 初始化TaobaoopenimtribeexpelAPIRequest对象
func NewTaobaoopenimtribeexpelRequest() *TaobaoopenimtribeexpelAPIRequest {
	return &TaobaoopenimtribeexpelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenimtribeexpelAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.expel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenimtribeexpelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenimtribeexpelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoopenimtribeexpelAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoopenimtribeexpelAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// SetTribeId is TribeId Setter
// 群id
func (r *TaobaoopenimtribeexpelAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// GetTribeId TribeId Getter
func (r TaobaoopenimtribeexpelAPIRequest) GetTribeId() int64 {
	return r._tribeId
}

// SetMember is Member Setter
// 用户信息
func (r *TaobaoopenimtribeexpelAPIRequest) SetMember(_member *OpenImUser) error {
	r._member = _member
	r.Set("member", _member)
	return nil
}

// GetMember Member Getter
func (r TaobaoopenimtribeexpelAPIRequest) GetMember() *OpenImUser {
	return r._member
}
