package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIRequest 门店无隐形消费签约 API请求
// alibaba.alihealth.dental.store.invisible.consume.update
//
// 门店无隐形消费签约
type AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIRequest struct {
	model.Params
	// 入参
	_store *DentalOuterStoreNicRequest
}

// NewAlibabaAlihealthDentalStoreInvisibleConsumeUpdateRequest 初始化AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIRequest对象
func NewAlibabaAlihealthDentalStoreInvisibleConsumeUpdateRequest() *AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIRequest {
	return &AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.dental.store.invisible.consume.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStore is Store Setter
// 入参
func (r *AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIRequest) SetStore(_store *DentalOuterStoreNicRequest) error {
	r._store = _store
	r.Set("store", _store)
	return nil
}

// GetStore Store Getter
func (r AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIRequest) GetStore() *DentalOuterStoreNicRequest {
	return r._store
}
