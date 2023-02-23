package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthReserveDentalUnbinditemAPIRequest 解绑商品信息 API请求
// alibaba.alihealth.reserve.dental.unbinditem
//
// 绑定门店信息，商品信息
type AlibabaAlihealthReserveDentalUnbinditemAPIRequest struct {
	model.Params
	// 服务商门店id
	_spStoreId string
	// 服务商商品id
	_spItemId string
}

// NewAlibabaAlihealthReserveDentalUnbinditemRequest 初始化AlibabaAlihealthReserveDentalUnbinditemAPIRequest对象
func NewAlibabaAlihealthReserveDentalUnbinditemRequest() *AlibabaAlihealthReserveDentalUnbinditemAPIRequest {
	return &AlibabaAlihealthReserveDentalUnbinditemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthReserveDentalUnbinditemAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.reserve.dental.unbinditem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthReserveDentalUnbinditemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthReserveDentalUnbinditemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSpStoreId is SpStoreId Setter
// 服务商门店id
func (r *AlibabaAlihealthReserveDentalUnbinditemAPIRequest) SetSpStoreId(_spStoreId string) error {
	r._spStoreId = _spStoreId
	r.Set("sp_store_id", _spStoreId)
	return nil
}

// GetSpStoreId SpStoreId Getter
func (r AlibabaAlihealthReserveDentalUnbinditemAPIRequest) GetSpStoreId() string {
	return r._spStoreId
}

// SetSpItemId is SpItemId Setter
// 服务商商品id
func (r *AlibabaAlihealthReserveDentalUnbinditemAPIRequest) SetSpItemId(_spItemId string) error {
	r._spItemId = _spItemId
	r.Set("sp_item_id", _spItemId)
	return nil
}

// GetSpItemId SpItemId Getter
func (r AlibabaAlihealthReserveDentalUnbinditemAPIRequest) GetSpItemId() string {
	return r._spItemId
}
