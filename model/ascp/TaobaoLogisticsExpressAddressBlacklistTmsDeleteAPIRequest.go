package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest 上门取退可揽范围黑名单删除接口 API请求
// taobao.logistics.express.address.blacklist.tms.delete
//
// 上门取退可揽范围黑名单删除接口
type TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest struct {
	model.Params
	// 上门取退可揽范围黑名单删除
	_addressBlacklistDeleteRequest *AddressBlacklistDeleteRequest
}

// NewTaobaoLogisticsExpressAddressBlacklistTmsDeleteRequest 初始化TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest对象
func NewTaobaoLogisticsExpressAddressBlacklistTmsDeleteRequest() *TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest {
	return &TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest) Reset() {
	r._addressBlacklistDeleteRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.address.blacklist.tms.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAddressBlacklistDeleteRequest is AddressBlacklistDeleteRequest Setter
// 上门取退可揽范围黑名单删除
func (r *TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest) SetAddressBlacklistDeleteRequest(_addressBlacklistDeleteRequest *AddressBlacklistDeleteRequest) error {
	r._addressBlacklistDeleteRequest = _addressBlacklistDeleteRequest
	r.Set("address_blacklist_delete_request", _addressBlacklistDeleteRequest)
	return nil
}

// GetAddressBlacklistDeleteRequest AddressBlacklistDeleteRequest Getter
func (r TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest) GetAddressBlacklistDeleteRequest() *AddressBlacklistDeleteRequest {
	return r._addressBlacklistDeleteRequest
}

var poolTaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsExpressAddressBlacklistTmsDeleteRequest()
	},
}

// GetTaobaoLogisticsExpressAddressBlacklistTmsDeleteRequest 从 sync.Pool 获取 TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest
func GetTaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest() *TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest {
	return poolTaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest.Get().(*TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest)
}

// ReleaseTaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest 将 TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest(v *TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest.Put(v)
}
