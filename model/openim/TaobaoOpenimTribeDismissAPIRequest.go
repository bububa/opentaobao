package openim

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeDismissAPIRequest OPENIM群解散 API请求
// taobao.openim.tribe.dismiss
//
// OPENIM群解散
type TaobaoOpenimTribeDismissAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tribeId int64
}

// NewTaobaoOpenimTribeDismissRequest 初始化TaobaoOpenimTribeDismissAPIRequest对象
func NewTaobaoOpenimTribeDismissRequest() *TaobaoOpenimTribeDismissAPIRequest {
	return &TaobaoOpenimTribeDismissAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenimTribeDismissAPIRequest) Reset() {
	r._user = nil
	r._tribeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeDismissAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.dismiss"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeDismissAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenimTribeDismissAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoOpenimTribeDismissAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoOpenimTribeDismissAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// SetTribeId is TribeId Setter
// 群id
func (r *TaobaoOpenimTribeDismissAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// GetTribeId TribeId Getter
func (r TaobaoOpenimTribeDismissAPIRequest) GetTribeId() int64 {
	return r._tribeId
}

var poolTaobaoOpenimTribeDismissAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenimTribeDismissRequest()
	},
}

// GetTaobaoOpenimTribeDismissRequest 从 sync.Pool 获取 TaobaoOpenimTribeDismissAPIRequest
func GetTaobaoOpenimTribeDismissAPIRequest() *TaobaoOpenimTribeDismissAPIRequest {
	return poolTaobaoOpenimTribeDismissAPIRequest.Get().(*TaobaoOpenimTribeDismissAPIRequest)
}

// ReleaseTaobaoOpenimTribeDismissAPIRequest 将 TaobaoOpenimTribeDismissAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenimTribeDismissAPIRequest(v *TaobaoOpenimTribeDismissAPIRequest) {
	v.Reset()
	poolTaobaoOpenimTribeDismissAPIRequest.Put(v)
}
