package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdentalstoreinsertorupdateAPIRequest ISV新增/修改口腔门店 API请求
// alibaba.alihealth.dental.store.insertorupdate
//
// ISV新增/修改口腔门店
type AlibabaalihealthdentalstoreinsertorupdateAPIRequest struct {
	model.Params
	// 门店
	_store *DentalOuterStoreRequest
}

// NewAlibabaalihealthdentalstoreinsertorupdateRequest 初始化AlibabaalihealthdentalstoreinsertorupdateAPIRequest对象
func NewAlibabaalihealthdentalstoreinsertorupdateRequest() *AlibabaalihealthdentalstoreinsertorupdateAPIRequest {
	return &AlibabaalihealthdentalstoreinsertorupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdentalstoreinsertorupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.dental.store.insertorupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdentalstoreinsertorupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdentalstoreinsertorupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStore is Store Setter
// 门店
func (r *AlibabaalihealthdentalstoreinsertorupdateAPIRequest) SetStore(_store *DentalOuterStoreRequest) error {
	r._store = _store
	r.Set("store", _store)
	return nil
}

// GetStore Store Getter
func (r AlibabaalihealthdentalstoreinsertorupdateAPIRequest) GetStore() *DentalOuterStoreRequest {
	return r._store
}
