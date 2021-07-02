package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeExpelAPIRequest OPENIM群踢出成员 API请求
// taobao.openim.tribe.expel
//
// OPENIM群踢出成员
type TaobaoOpenimTribeExpelAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tribeId int64
	// 用户信息
	_member *OpenImUser
}

// NewTaobaoOpenimTribeExpelRequest 初始化TaobaoOpenimTribeExpelAPIRequest对象
func NewTaobaoOpenimTribeExpelRequest() *TaobaoOpenimTribeExpelAPIRequest {
	return &TaobaoOpenimTribeExpelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeExpelAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.expel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeExpelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is User Setter
// 用户信息
func (r *TaobaoOpenimTribeExpelAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// Get User Getter
func (r TaobaoOpenimTribeExpelAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// Set is TribeId Setter
// 群id
func (r *TaobaoOpenimTribeExpelAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// Get TribeId Getter
func (r TaobaoOpenimTribeExpelAPIRequest) GetTribeId() int64 {
	return r._tribeId
}

// Set is Member Setter
// 用户信息
func (r *TaobaoOpenimTribeExpelAPIRequest) SetMember(_member *OpenImUser) error {
	r._member = _member
	r.Set("member", _member)
	return nil
}

// Get Member Getter
func (r TaobaoOpenimTribeExpelAPIRequest) GetMember() *OpenImUser {
	return r._member
}
