package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimTribeDismissAPIRequest
OPENIM群解散 API请求
taobao.openim.tribe.dismiss

OPENIM群解散 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeDismissAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.dismiss"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeDismissAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is User Setter
// 用户信息
func (r *TaobaoOpenimTribeDismissAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// Get User Getter
func (r TaobaoOpenimTribeDismissAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// Set is TribeId Setter
// 群id
func (r *TaobaoOpenimTribeDismissAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// Get TribeId Getter
func (r TaobaoOpenimTribeDismissAPIRequest) GetTribeId() int64 {
	return r._tribeId
}
