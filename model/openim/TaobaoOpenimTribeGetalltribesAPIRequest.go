package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimtribegetalltribesAPIRequest 获取用户群列表 API请求
// taobao.openim.tribe.getalltribes
//
// OPENIM群服务获取用户群列表
type TaobaoopenimtribegetalltribesAPIRequest struct {
	model.Params
	// 群类型
	_tribeTypes []int64
	// 用户信息
	_user *OpenImUser
}

// NewTaobaoopenimtribegetalltribesRequest 初始化TaobaoopenimtribegetalltribesAPIRequest对象
func NewTaobaoopenimtribegetalltribesRequest() *TaobaoopenimtribegetalltribesAPIRequest {
	return &TaobaoopenimtribegetalltribesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenimtribegetalltribesAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.getalltribes"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenimtribegetalltribesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenimtribegetalltribesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTribeTypes is TribeTypes Setter
// 群类型
func (r *TaobaoopenimtribegetalltribesAPIRequest) SetTribeTypes(_tribeTypes []int64) error {
	r._tribeTypes = _tribeTypes
	r.Set("tribe_types", _tribeTypes)
	return nil
}

// GetTribeTypes TribeTypes Getter
func (r TaobaoopenimtribegetalltribesAPIRequest) GetTribeTypes() []int64 {
	return r._tribeTypes
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoopenimtribegetalltribesAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoopenimtribegetalltribesAPIRequest) GetUser() *OpenImUser {
	return r._user
}
