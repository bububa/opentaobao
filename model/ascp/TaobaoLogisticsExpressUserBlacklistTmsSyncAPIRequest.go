package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpressuserblacklisttmssyncAPIRequest 上门取退用户黑名单同步 API请求
// taobao.logistics.express.user.blacklist.tms.sync
//
// 上门取退用户黑名单同步
type TaobaologisticsexpressuserblacklisttmssyncAPIRequest struct {
	model.Params
	// 上门取退用户黑名单
	_userBlacklistRequest *UserBlacklistRequest
}

// NewTaobaologisticsexpressuserblacklisttmssyncRequest 初始化TaobaologisticsexpressuserblacklisttmssyncAPIRequest对象
func NewTaobaologisticsexpressuserblacklisttmssyncRequest() *TaobaologisticsexpressuserblacklisttmssyncAPIRequest {
	return &TaobaologisticsexpressuserblacklisttmssyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsexpressuserblacklisttmssyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.user.blacklist.tms.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsexpressuserblacklisttmssyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsexpressuserblacklisttmssyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserBlacklistRequest is UserBlacklistRequest Setter
// 上门取退用户黑名单
func (r *TaobaologisticsexpressuserblacklisttmssyncAPIRequest) SetUserBlacklistRequest(_userBlacklistRequest *UserBlacklistRequest) error {
	r._userBlacklistRequest = _userBlacklistRequest
	r.Set("user_blacklist_request", _userBlacklistRequest)
	return nil
}

// GetUserBlacklistRequest UserBlacklistRequest Getter
func (r TaobaologisticsexpressuserblacklisttmssyncAPIRequest) GetUserBlacklistRequest() *UserBlacklistRequest {
	return r._userBlacklistRequest
}
