package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpressaddressblacklisttmsasyncAPIRequest 上门取退可揽范围黑名单同步/更新 API请求
// taobao.logistics.express.address.blacklist.tms.async
//
// 上门取退可揽范围黑名单同步/更新
type TaobaologisticsexpressaddressblacklisttmsasyncAPIRequest struct {
	model.Params
	// 上门取退可揽范围黑名
	_addressBlacklistRequest *AddressBlacklistRequest
}

// NewTaobaologisticsexpressaddressblacklisttmsasyncRequest 初始化TaobaologisticsexpressaddressblacklisttmsasyncAPIRequest对象
func NewTaobaologisticsexpressaddressblacklisttmsasyncRequest() *TaobaologisticsexpressaddressblacklisttmsasyncAPIRequest {
	return &TaobaologisticsexpressaddressblacklisttmsasyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsexpressaddressblacklisttmsasyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.address.blacklist.tms.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsexpressaddressblacklisttmsasyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsexpressaddressblacklisttmsasyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAddressBlacklistRequest is AddressBlacklistRequest Setter
// 上门取退可揽范围黑名
func (r *TaobaologisticsexpressaddressblacklisttmsasyncAPIRequest) SetAddressBlacklistRequest(_addressBlacklistRequest *AddressBlacklistRequest) error {
	r._addressBlacklistRequest = _addressBlacklistRequest
	r.Set("address_blacklist_request", _addressBlacklistRequest)
	return nil
}

// GetAddressBlacklistRequest AddressBlacklistRequest Getter
func (r TaobaologisticsexpressaddressblacklisttmsasyncAPIRequest) GetAddressBlacklistRequest() *AddressBlacklistRequest {
	return r._addressBlacklistRequest
}
