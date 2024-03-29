package openim

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeCreateAPIRequest 创建群 API请求
// taobao.openim.tribe.create
//
// 创建一个openim的群
type TaobaoOpenimTribeCreateAPIRequest struct {
	model.Params
	// 创建群时候拉入群的成员tribe_type = 1（即为讨论组类型)时  该参数为必选tribe_type = 0 (即为普通群类型)时，改参数无效，可不填
	_members []OpenImUser
	// 群名称
	_tribeName string
	// 群公告
	_notice string
	// 用户信息
	_user *OpenImUser
	// 群类型有两种tribe_type = 0  普通群  普通群有管理员角色，对成员加入有权限控制tribe_type = 1  讨论组  讨论组没有管理员，不能解散
	_tribeType int64
}

// NewTaobaoOpenimTribeCreateRequest 初始化TaobaoOpenimTribeCreateAPIRequest对象
func NewTaobaoOpenimTribeCreateRequest() *TaobaoOpenimTribeCreateAPIRequest {
	return &TaobaoOpenimTribeCreateAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenimTribeCreateAPIRequest) Reset() {
	r._members = r._members[:0]
	r._tribeName = ""
	r._notice = ""
	r._user = nil
	r._tribeType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeCreateAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenimTribeCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMembers is Members Setter
// 创建群时候拉入群的成员tribe_type = 1（即为讨论组类型)时  该参数为必选tribe_type = 0 (即为普通群类型)时，改参数无效，可不填
func (r *TaobaoOpenimTribeCreateAPIRequest) SetMembers(_members []OpenImUser) error {
	r._members = _members
	r.Set("members", _members)
	return nil
}

// GetMembers Members Getter
func (r TaobaoOpenimTribeCreateAPIRequest) GetMembers() []OpenImUser {
	return r._members
}

// SetTribeName is TribeName Setter
// 群名称
func (r *TaobaoOpenimTribeCreateAPIRequest) SetTribeName(_tribeName string) error {
	r._tribeName = _tribeName
	r.Set("tribe_name", _tribeName)
	return nil
}

// GetTribeName TribeName Getter
func (r TaobaoOpenimTribeCreateAPIRequest) GetTribeName() string {
	return r._tribeName
}

// SetNotice is Notice Setter
// 群公告
func (r *TaobaoOpenimTribeCreateAPIRequest) SetNotice(_notice string) error {
	r._notice = _notice
	r.Set("notice", _notice)
	return nil
}

// GetNotice Notice Getter
func (r TaobaoOpenimTribeCreateAPIRequest) GetNotice() string {
	return r._notice
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoOpenimTribeCreateAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoOpenimTribeCreateAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// SetTribeType is TribeType Setter
// 群类型有两种tribe_type = 0  普通群  普通群有管理员角色，对成员加入有权限控制tribe_type = 1  讨论组  讨论组没有管理员，不能解散
func (r *TaobaoOpenimTribeCreateAPIRequest) SetTribeType(_tribeType int64) error {
	r._tribeType = _tribeType
	r.Set("tribe_type", _tribeType)
	return nil
}

// GetTribeType TribeType Getter
func (r TaobaoOpenimTribeCreateAPIRequest) GetTribeType() int64 {
	return r._tribeType
}

var poolTaobaoOpenimTribeCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenimTribeCreateRequest()
	},
}

// GetTaobaoOpenimTribeCreateRequest 从 sync.Pool 获取 TaobaoOpenimTribeCreateAPIRequest
func GetTaobaoOpenimTribeCreateAPIRequest() *TaobaoOpenimTribeCreateAPIRequest {
	return poolTaobaoOpenimTribeCreateAPIRequest.Get().(*TaobaoOpenimTribeCreateAPIRequest)
}

// ReleaseTaobaoOpenimTribeCreateAPIRequest 将 TaobaoOpenimTribeCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenimTribeCreateAPIRequest(v *TaobaoOpenimTribeCreateAPIRequest) {
	v.Reset()
	poolTaobaoOpenimTribeCreateAPIRequest.Put(v)
}
