package openim

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenimTribeUnsetmanagerAPIRequest) Reset() {
	r._user = nil
	r._tid = 0
	r._member = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeUnsetmanagerAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.unsetmanager"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeUnsetmanagerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenimTribeUnsetmanagerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoOpenimTribeUnsetmanagerAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoOpenimTribeUnsetmanagerAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// SetTid is Tid Setter
// 群id
func (r *TaobaoOpenimTribeUnsetmanagerAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoOpenimTribeUnsetmanagerAPIRequest) GetTid() int64 {
	return r._tid
}

// SetMember is Member Setter
// 用户信息
func (r *TaobaoOpenimTribeUnsetmanagerAPIRequest) SetMember(_member *OpenImUser) error {
	r._member = _member
	r.Set("member", _member)
	return nil
}

// GetMember Member Getter
func (r TaobaoOpenimTribeUnsetmanagerAPIRequest) GetMember() *OpenImUser {
	return r._member
}

var poolTaobaoOpenimTribeUnsetmanagerAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenimTribeUnsetmanagerRequest()
	},
}

// GetTaobaoOpenimTribeUnsetmanagerRequest 从 sync.Pool 获取 TaobaoOpenimTribeUnsetmanagerAPIRequest
func GetTaobaoOpenimTribeUnsetmanagerAPIRequest() *TaobaoOpenimTribeUnsetmanagerAPIRequest {
	return poolTaobaoOpenimTribeUnsetmanagerAPIRequest.Get().(*TaobaoOpenimTribeUnsetmanagerAPIRequest)
}

// ReleaseTaobaoOpenimTribeUnsetmanagerAPIRequest 将 TaobaoOpenimTribeUnsetmanagerAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenimTribeUnsetmanagerAPIRequest(v *TaobaoOpenimTribeUnsetmanagerAPIRequest) {
	v.Reset()
	poolTaobaoOpenimTribeUnsetmanagerAPIRequest.Put(v)
}
