package openim

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeGettribeinfoAPIRequest 获取群信息 API请求
// taobao.openim.tribe.gettribeinfo
//
// 获取群信息
type TaobaoOpenimTribeGettribeinfoAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群ID
	_tribeId int64
}

// NewTaobaoOpenimTribeGettribeinfoRequest 初始化TaobaoOpenimTribeGettribeinfoAPIRequest对象
func NewTaobaoOpenimTribeGettribeinfoRequest() *TaobaoOpenimTribeGettribeinfoAPIRequest {
	return &TaobaoOpenimTribeGettribeinfoAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenimTribeGettribeinfoAPIRequest) Reset() {
	r._user = nil
	r._tribeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeGettribeinfoAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.gettribeinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeGettribeinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenimTribeGettribeinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoOpenimTribeGettribeinfoAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoOpenimTribeGettribeinfoAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// SetTribeId is TribeId Setter
// 群ID
func (r *TaobaoOpenimTribeGettribeinfoAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// GetTribeId TribeId Getter
func (r TaobaoOpenimTribeGettribeinfoAPIRequest) GetTribeId() int64 {
	return r._tribeId
}

var poolTaobaoOpenimTribeGettribeinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenimTribeGettribeinfoRequest()
	},
}

// GetTaobaoOpenimTribeGettribeinfoRequest 从 sync.Pool 获取 TaobaoOpenimTribeGettribeinfoAPIRequest
func GetTaobaoOpenimTribeGettribeinfoAPIRequest() *TaobaoOpenimTribeGettribeinfoAPIRequest {
	return poolTaobaoOpenimTribeGettribeinfoAPIRequest.Get().(*TaobaoOpenimTribeGettribeinfoAPIRequest)
}

// ReleaseTaobaoOpenimTribeGettribeinfoAPIRequest 将 TaobaoOpenimTribeGettribeinfoAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenimTribeGettribeinfoAPIRequest(v *TaobaoOpenimTribeGettribeinfoAPIRequest) {
	v.Reset()
	poolTaobaoOpenimTribeGettribeinfoAPIRequest.Put(v)
}
