package openim

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeGettribeinfoAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.gettribeinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeGettribeinfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
