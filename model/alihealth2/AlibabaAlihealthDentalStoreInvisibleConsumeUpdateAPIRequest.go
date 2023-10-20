package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdentalstoreinvisibleconsumeupdateAPIRequest 门店无隐形消费签约 API请求
// alibaba.alihealth.dental.store.invisible.consume.update
//
// 门店无隐形消费签约
type AlibabaalihealthdentalstoreinvisibleconsumeupdateAPIRequest struct {
	model.Params
	// 入参
	_store *DentalOuterStoreNicRequest
}

// NewAlibabaalihealthdentalstoreinvisibleconsumeupdateRequest 初始化AlibabaalihealthdentalstoreinvisibleconsumeupdateAPIRequest对象
func NewAlibabaalihealthdentalstoreinvisibleconsumeupdateRequest() *AlibabaalihealthdentalstoreinvisibleconsumeupdateAPIRequest {
	return &AlibabaalihealthdentalstoreinvisibleconsumeupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdentalstoreinvisibleconsumeupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.dental.store.invisible.consume.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdentalstoreinvisibleconsumeupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdentalstoreinvisibleconsumeupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStore is Store Setter
// 入参
func (r *AlibabaalihealthdentalstoreinvisibleconsumeupdateAPIRequest) SetStore(_store *DentalOuterStoreNicRequest) error {
	r._store = _store
	r.Set("store", _store)
	return nil
}

// GetStore Store Getter
func (r AlibabaalihealthdentalstoreinvisibleconsumeupdateAPIRequest) GetStore() *DentalOuterStoreNicRequest {
	return r._store
}
