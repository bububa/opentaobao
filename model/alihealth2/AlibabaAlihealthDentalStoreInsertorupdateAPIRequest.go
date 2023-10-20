package alihealth2

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDentalStoreInsertorupdateAPIRequest) Reset() {
	r._store = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalStoreInsertorupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.dental.store.insertorupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalStoreInsertorupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDentalStoreInsertorupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihealthDentalStoreInsertorupdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDentalStoreInsertorupdateRequest()
	},
}

// GetAlibabaAlihealthDentalStoreInsertorupdateRequest 从 sync.Pool 获取 AlibabaAlihealthDentalStoreInsertorupdateAPIRequest
func GetAlibabaAlihealthDentalStoreInsertorupdateAPIRequest() *AlibabaAlihealthDentalStoreInsertorupdateAPIRequest {
	return poolAlibabaAlihealthDentalStoreInsertorupdateAPIRequest.Get().(*AlibabaAlihealthDentalStoreInsertorupdateAPIRequest)
}

// ReleaseAlibabaAlihealthDentalStoreInsertorupdateAPIRequest 将 AlibabaAlihealthDentalStoreInsertorupdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDentalStoreInsertorupdateAPIRequest(v *AlibabaAlihealthDentalStoreInsertorupdateAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDentalStoreInsertorupdateAPIRequest.Put(v)
}
