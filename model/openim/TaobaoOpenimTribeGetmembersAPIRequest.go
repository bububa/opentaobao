package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeGetmembersAPIRequest OPENIM群成员获取 API请求
// taobao.openim.tribe.getmembers
//
// OPENIM群成员获取
type TaobaoOpenimTribeGetmembersAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tribeId int64
}

// NewTaobaoOpenimTribeGetmembersRequest 初始化TaobaoOpenimTribeGetmembersAPIRequest对象
func NewTaobaoOpenimTribeGetmembersRequest() *TaobaoOpenimTribeGetmembersAPIRequest {
	return &TaobaoOpenimTribeGetmembersAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeGetmembersAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.getmembers"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeGetmembersAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoOpenimTribeGetmembersAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoOpenimTribeGetmembersAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// SetTribeId is TribeId Setter
// 群id
func (r *TaobaoOpenimTribeGetmembersAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// GetTribeId TribeId Getter
func (r TaobaoOpenimTribeGetmembersAPIRequest) GetTribeId() int64 {
	return r._tribeId
}
