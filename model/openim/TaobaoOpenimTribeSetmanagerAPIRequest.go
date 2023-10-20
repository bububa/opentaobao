package openim

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeSetmanagerAPIRequest OPENIM群设置管理员 API请求
// taobao.openim.tribe.setmanager
//
// OPENIM群设置管理员
type TaobaoOpenimTribeSetmanagerAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tid int64
	// 用户信息
	_member *OpenImUser
}

// NewTaobaoOpenimTribeSetmanagerRequest 初始化TaobaoOpenimTribeSetmanagerAPIRequest对象
func NewTaobaoOpenimTribeSetmanagerRequest() *TaobaoOpenimTribeSetmanagerAPIRequest {
	return &TaobaoOpenimTribeSetmanagerAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenimTribeSetmanagerAPIRequest) Reset() {
	r._user = nil
	r._tid = 0
	r._member = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeSetmanagerAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.setmanager"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeSetmanagerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenimTribeSetmanagerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoOpenimTribeSetmanagerAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoOpenimTribeSetmanagerAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// SetTid is Tid Setter
// 群id
func (r *TaobaoOpenimTribeSetmanagerAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoOpenimTribeSetmanagerAPIRequest) GetTid() int64 {
	return r._tid
}

// SetMember is Member Setter
// 用户信息
func (r *TaobaoOpenimTribeSetmanagerAPIRequest) SetMember(_member *OpenImUser) error {
	r._member = _member
	r.Set("member", _member)
	return nil
}

// GetMember Member Getter
func (r TaobaoOpenimTribeSetmanagerAPIRequest) GetMember() *OpenImUser {
	return r._member
}

var poolTaobaoOpenimTribeSetmanagerAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenimTribeSetmanagerRequest()
	},
}

// GetTaobaoOpenimTribeSetmanagerRequest 从 sync.Pool 获取 TaobaoOpenimTribeSetmanagerAPIRequest
func GetTaobaoOpenimTribeSetmanagerAPIRequest() *TaobaoOpenimTribeSetmanagerAPIRequest {
	return poolTaobaoOpenimTribeSetmanagerAPIRequest.Get().(*TaobaoOpenimTribeSetmanagerAPIRequest)
}

// ReleaseTaobaoOpenimTribeSetmanagerAPIRequest 将 TaobaoOpenimTribeSetmanagerAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenimTribeSetmanagerAPIRequest(v *TaobaoOpenimTribeSetmanagerAPIRequest) {
	v.Reset()
	poolTaobaoOpenimTribeSetmanagerAPIRequest.Put(v)
}
