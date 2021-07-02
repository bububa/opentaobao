package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeGetalltribesAPIRequest 获取用户群列表 API请求
// taobao.openim.tribe.getalltribes
//
// OPENIM群服务获取用户群列表
type TaobaoOpenimTribeGetalltribesAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群类型
	_tribeTypes []int64
}

// NewTaobaoOpenimTribeGetalltribesRequest 初始化TaobaoOpenimTribeGetalltribesAPIRequest对象
func NewTaobaoOpenimTribeGetalltribesRequest() *TaobaoOpenimTribeGetalltribesAPIRequest {
	return &TaobaoOpenimTribeGetalltribesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeGetalltribesAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.getalltribes"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeGetalltribesAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is User Setter
// 用户信息
func (r *TaobaoOpenimTribeGetalltribesAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// Get User Getter
func (r TaobaoOpenimTribeGetalltribesAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// Set is TribeTypes Setter
// 群类型
func (r *TaobaoOpenimTribeGetalltribesAPIRequest) SetTribeTypes(_tribeTypes []int64) error {
	r._tribeTypes = _tribeTypes
	r.Set("tribe_types", _tribeTypes)
	return nil
}

// Get TribeTypes Getter
func (r TaobaoOpenimTribeGetalltribesAPIRequest) GetTribeTypes() []int64 {
	return r._tribeTypes
}
