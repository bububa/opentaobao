package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDentalStoreInsertorupdateAPIRequest ISV新增/修改口腔门店 API请求
// alibaba.alihealth.dental.store.insertorupdate
//
// ISV新增/修改口腔门店
type AlibabaAlihealthDentalStoreInsertorupdateAPIRequest struct {
	model.Params
	// 门店
	_store *DentalOuterStoreRequest
}

// NewAlibabaAlihealthDentalStoreInsertorupdateRequest 初始化AlibabaAlihealthDentalStoreInsertorupdateAPIRequest对象
func NewAlibabaAlihealthDentalStoreInsertorupdateRequest() *AlibabaAlihealthDentalStoreInsertorupdateAPIRequest {
	return &AlibabaAlihealthDentalStoreInsertorupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalStoreInsertorupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.dental.store.insertorupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalStoreInsertorupdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStore is Store Setter
// 门店
func (r *AlibabaAlihealthDentalStoreInsertorupdateAPIRequest) SetStore(_store *DentalOuterStoreRequest) error {
	r._store = _store
	r.Set("store", _store)
	return nil
}

// GetStore Store Getter
func (r AlibabaAlihealthDentalStoreInsertorupdateAPIRequest) GetStore() *DentalOuterStoreRequest {
	return r._store
}
