package openim

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeInviteAPIRequest OPENIM群邀请加入 API请求
// taobao.openim.tribe.invite
//
// OPENIM群邀请加入接口
type TaobaoOpenimTribeInviteAPIRequest struct {
	model.Params
	// 用户信息
	_members []OpenImUser
	// 群id
	_tribeId int64
	// 用户信息
	_user *OpenImUser
}

// NewTaobaoOpenimTribeInviteRequest 初始化TaobaoOpenimTribeInviteAPIRequest对象
func NewTaobaoOpenimTribeInviteRequest() *TaobaoOpenimTribeInviteAPIRequest {
	return &TaobaoOpenimTribeInviteAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenimTribeInviteAPIRequest) Reset() {
	r._members = r._members[:0]
	r._tribeId = 0
	r._user = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeInviteAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.invite"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeInviteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenimTribeInviteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMembers is Members Setter
// 用户信息
func (r *TaobaoOpenimTribeInviteAPIRequest) SetMembers(_members []OpenImUser) error {
	r._members = _members
	r.Set("members", _members)
	return nil
}

// GetMembers Members Getter
func (r TaobaoOpenimTribeInviteAPIRequest) GetMembers() []OpenImUser {
	return r._members
}

// SetTribeId is TribeId Setter
// 群id
func (r *TaobaoOpenimTribeInviteAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// GetTribeId TribeId Getter
func (r TaobaoOpenimTribeInviteAPIRequest) GetTribeId() int64 {
	return r._tribeId
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoOpenimTribeInviteAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoOpenimTribeInviteAPIRequest) GetUser() *OpenImUser {
	return r._user
}

var poolTaobaoOpenimTribeInviteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenimTribeInviteRequest()
	},
}

// GetTaobaoOpenimTribeInviteRequest 从 sync.Pool 获取 TaobaoOpenimTribeInviteAPIRequest
func GetTaobaoOpenimTribeInviteAPIRequest() *TaobaoOpenimTribeInviteAPIRequest {
	return poolTaobaoOpenimTribeInviteAPIRequest.Get().(*TaobaoOpenimTribeInviteAPIRequest)
}

// ReleaseTaobaoOpenimTribeInviteAPIRequest 将 TaobaoOpenimTribeInviteAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenimTribeInviteAPIRequest(v *TaobaoOpenimTribeInviteAPIRequest) {
	v.Reset()
	poolTaobaoOpenimTribeInviteAPIRequest.Put(v)
}
