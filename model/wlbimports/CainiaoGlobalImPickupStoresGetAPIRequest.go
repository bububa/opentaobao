package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupStoresGetAPIRequest 首公里揽收-集货仓列表查询 API请求
// cainiao.global.im.pickup.stores.get
//
// 首公里揽收-集货仓列表查询
type CainiaoGlobalImPickupStoresGetAPIRequest struct {
	model.Params
	// 请求体
	_transferstoreQueryRequest *TransferstoreQueryRequest
}

// NewCainiaoGlobalImPickupStoresGetRequest 初始化CainiaoGlobalImPickupStoresGetAPIRequest对象
func NewCainiaoGlobalImPickupStoresGetRequest() *CainiaoGlobalImPickupStoresGetAPIRequest {
	return &CainiaoGlobalImPickupStoresGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalImPickupStoresGetAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.stores.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalImPickupStoresGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTransferstoreQueryRequest is TransferstoreQueryRequest Setter
// 请求体
func (r *CainiaoGlobalImPickupStoresGetAPIRequest) SetTransferstoreQueryRequest(_transferstoreQueryRequest *TransferstoreQueryRequest) error {
	r._transferstoreQueryRequest = _transferstoreQueryRequest
	r.Set("transferstore_query_request", _transferstoreQueryRequest)
	return nil
}

// GetTransferstoreQueryRequest TransferstoreQueryRequest Getter
func (r CainiaoGlobalImPickupStoresGetAPIRequest) GetTransferstoreQueryRequest() *TransferstoreQueryRequest {
	return r._transferstoreQueryRequest
}
