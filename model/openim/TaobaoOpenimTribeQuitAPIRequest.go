package openim

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeQuitAPIRequest OPENIM群成员退出 API请求
// taobao.openim.tribe.quit
//
// OPENIM群成员退出
type TaobaoOpenimTribeQuitAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tribeId int64
}

// NewTaobaoOpenimTribeQuitRequest 初始化TaobaoOpenimTribeQuitAPIRequest对象
func NewTaobaoOpenimTribeQuitRequest() *TaobaoOpenimTribeQuitAPIRequest {
	return &TaobaoOpenimTribeQuitAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenimTribeQuitAPIRequest) Reset() {
	r._user = nil
	r._tribeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeQuitAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.quit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeQuitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenimTribeQuitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoOpenimTribeQuitAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoOpenimTribeQuitAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// SetTribeId is TribeId Setter
// 群id
func (r *TaobaoOpenimTribeQuitAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// GetTribeId TribeId Getter
func (r TaobaoOpenimTribeQuitAPIRequest) GetTribeId() int64 {
	return r._tribeId
}

var poolTaobaoOpenimTribeQuitAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenimTribeQuitRequest()
	},
}

// GetTaobaoOpenimTribeQuitRequest 从 sync.Pool 获取 TaobaoOpenimTribeQuitAPIRequest
func GetTaobaoOpenimTribeQuitAPIRequest() *TaobaoOpenimTribeQuitAPIRequest {
	return poolTaobaoOpenimTribeQuitAPIRequest.Get().(*TaobaoOpenimTribeQuitAPIRequest)
}

// ReleaseTaobaoOpenimTribeQuitAPIRequest 将 TaobaoOpenimTribeQuitAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenimTribeQuitAPIRequest(v *TaobaoOpenimTribeQuitAPIRequest) {
	v.Reset()
	poolTaobaoOpenimTribeQuitAPIRequest.Put(v)
}
