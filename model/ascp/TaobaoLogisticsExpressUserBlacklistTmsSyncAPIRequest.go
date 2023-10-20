package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressUserBlacklistTmsSyncAPIRequest 上门取退用户黑名单同步 API请求
// taobao.logistics.express.user.blacklist.tms.sync
//
// 上门取退用户黑名单同步
type TaobaoLogisticsExpressUserBlacklistTmsSyncAPIRequest struct {
	model.Params
	// 上门取退用户黑名单
	_userBlacklistRequest *UserBlacklistRequest
}

// NewTaobaoLogisticsExpressUserBlacklistTmsSyncRequest 初始化TaobaoLogisticsExpressUserBlacklistTmsSyncAPIRequest对象
func NewTaobaoLogisticsExpressUserBlacklistTmsSyncRequest() *TaobaoLogisticsExpressUserBlacklistTmsSyncAPIRequest {
	return &TaobaoLogisticsExpressUserBlacklistTmsSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsExpressUserBlacklistTmsSyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.user.blacklist.tms.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsExpressUserBlacklistTmsSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsExpressUserBlacklistTmsSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserBlacklistRequest is UserBlacklistRequest Setter
// 上门取退用户黑名单
func (r *TaobaoLogisticsExpressUserBlacklistTmsSyncAPIRequest) SetUserBlacklistRequest(_userBlacklistRequest *UserBlacklistRequest) error {
	r._userBlacklistRequest = _userBlacklistRequest
	r.Set("user_blacklist_request", _userBlacklistRequest)
	return nil
}

// GetUserBlacklistRequest UserBlacklistRequest Getter
func (r TaobaoLogisticsExpressUserBlacklistTmsSyncAPIRequest) GetUserBlacklistRequest() *UserBlacklistRequest {
	return r._userBlacklistRequest
}
