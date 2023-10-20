package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest 上门取退可揽范围黑名单同步/更新 API请求
// taobao.logistics.express.address.blacklist.tms.async
//
// 上门取退可揽范围黑名单同步/更新
type TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest struct {
	model.Params
	// 上门取退可揽范围黑名
	_addressBlacklistRequest *AddressBlacklistRequest
}

// NewTaobaoLogisticsExpressAddressBlacklistTmsAsyncRequest 初始化TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest对象
func NewTaobaoLogisticsExpressAddressBlacklistTmsAsyncRequest() *TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest {
	return &TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest) Reset() {
	r._addressBlacklistRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.address.blacklist.tms.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAddressBlacklistRequest is AddressBlacklistRequest Setter
// 上门取退可揽范围黑名
func (r *TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest) SetAddressBlacklistRequest(_addressBlacklistRequest *AddressBlacklistRequest) error {
	r._addressBlacklistRequest = _addressBlacklistRequest
	r.Set("address_blacklist_request", _addressBlacklistRequest)
	return nil
}

// GetAddressBlacklistRequest AddressBlacklistRequest Getter
func (r TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest) GetAddressBlacklistRequest() *AddressBlacklistRequest {
	return r._addressBlacklistRequest
}

var poolTaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsExpressAddressBlacklistTmsAsyncRequest()
	},
}

// GetTaobaoLogisticsExpressAddressBlacklistTmsAsyncRequest 从 sync.Pool 获取 TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest
func GetTaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest() *TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest {
	return poolTaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest.Get().(*TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest)
}

// ReleaseTaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest 将 TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest(v *TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest.Put(v)
}
