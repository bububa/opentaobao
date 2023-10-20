package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpressaddressblacklisttmsdeleteAPIRequest 上门取退可揽范围黑名单删除接口 API请求
// taobao.logistics.express.address.blacklist.tms.delete
//
// 上门取退可揽范围黑名单删除接口
type TaobaologisticsexpressaddressblacklisttmsdeleteAPIRequest struct {
	model.Params
	// 上门取退可揽范围黑名单删除
	_addressBlacklistDeleteRequest *AddressBlacklistDeleteRequest
}

// NewTaobaologisticsexpressaddressblacklisttmsdeleteRequest 初始化TaobaologisticsexpressaddressblacklisttmsdeleteAPIRequest对象
func NewTaobaologisticsexpressaddressblacklisttmsdeleteRequest() *TaobaologisticsexpressaddressblacklisttmsdeleteAPIRequest {
	return &TaobaologisticsexpressaddressblacklisttmsdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsexpressaddressblacklisttmsdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.address.blacklist.tms.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsexpressaddressblacklisttmsdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsexpressaddressblacklisttmsdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAddressBlacklistDeleteRequest is AddressBlacklistDeleteRequest Setter
// 上门取退可揽范围黑名单删除
func (r *TaobaologisticsexpressaddressblacklisttmsdeleteAPIRequest) SetAddressBlacklistDeleteRequest(_addressBlacklistDeleteRequest *AddressBlacklistDeleteRequest) error {
	r._addressBlacklistDeleteRequest = _addressBlacklistDeleteRequest
	r.Set("address_blacklist_delete_request", _addressBlacklistDeleteRequest)
	return nil
}

// GetAddressBlacklistDeleteRequest AddressBlacklistDeleteRequest Getter
func (r TaobaologisticsexpressaddressblacklisttmsdeleteAPIRequest) GetAddressBlacklistDeleteRequest() *AddressBlacklistDeleteRequest {
	return r._addressBlacklistDeleteRequest
}
