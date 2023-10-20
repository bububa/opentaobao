package openim

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenimTribeExpelAPIRequest) Reset() {
	r._user = nil
	r._tribeId = 0
	r._member = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeExpelAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.expel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeExpelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenimTribeExpelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoOpenimTribeExpelAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoOpenimTribeExpelAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// SetTribeId is TribeId Setter
// 群id
func (r *TaobaoOpenimTribeExpelAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// GetTribeId TribeId Getter
func (r TaobaoOpenimTribeExpelAPIRequest) GetTribeId() int64 {
	return r._tribeId
}

// SetMember is Member Setter
// 用户信息
func (r *TaobaoOpenimTribeExpelAPIRequest) SetMember(_member *OpenImUser) error {
	r._member = _member
	r.Set("member", _member)
	return nil
}

// GetMember Member Getter
func (r TaobaoOpenimTribeExpelAPIRequest) GetMember() *OpenImUser {
	return r._member
}

var poolTaobaoOpenimTribeExpelAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenimTribeExpelRequest()
	},
}

// GetTaobaoOpenimTribeExpelRequest 从 sync.Pool 获取 TaobaoOpenimTribeExpelAPIRequest
func GetTaobaoOpenimTribeExpelAPIRequest() *TaobaoOpenimTribeExpelAPIRequest {
	return poolTaobaoOpenimTribeExpelAPIRequest.Get().(*TaobaoOpenimTribeExpelAPIRequest)
}

// ReleaseTaobaoOpenimTribeExpelAPIRequest 将 TaobaoOpenimTribeExpelAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenimTribeExpelAPIRequest(v *TaobaoOpenimTribeExpelAPIRequest) {
	v.Reset()
	poolTaobaoOpenimTribeExpelAPIRequest.Put(v)
}
