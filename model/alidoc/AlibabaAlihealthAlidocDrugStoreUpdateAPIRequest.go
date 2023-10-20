package alidoc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthalidocdrugstoreupdateAPIRequest 更新药店 API请求
// alibaba.alihealth.alidoc.drug.store.update
//
// 药店信息更新接口
type AlibabaalihealthalidocdrugstoreupdateAPIRequest struct {
	model.Params
	// 更新对象
	_drugStoreUpdateTopRequest *DrugStoreUpdateTopRequest
}

// NewAlibabaalihealthalidocdrugstoreupdateRequest 初始化AlibabaalihealthalidocdrugstoreupdateAPIRequest对象
func NewAlibabaalihealthalidocdrugstoreupdateRequest() *AlibabaalihealthalidocdrugstoreupdateAPIRequest {
	return &AlibabaalihealthalidocdrugstoreupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthalidocdrugstoreupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.alidoc.drug.store.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthalidocdrugstoreupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthalidocdrugstoreupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDrugStoreUpdateTopRequest is DrugStoreUpdateTopRequest Setter
// 更新对象
func (r *AlibabaalihealthalidocdrugstoreupdateAPIRequest) SetDrugStoreUpdateTopRequest(_drugStoreUpdateTopRequest *DrugStoreUpdateTopRequest) error {
	r._drugStoreUpdateTopRequest = _drugStoreUpdateTopRequest
	r.Set("drug_store_update_top_request", _drugStoreUpdateTopRequest)
	return nil
}

// GetDrugStoreUpdateTopRequest DrugStoreUpdateTopRequest Getter
func (r AlibabaalihealthalidocdrugstoreupdateAPIRequest) GetDrugStoreUpdateTopRequest() *DrugStoreUpdateTopRequest {
	return r._drugStoreUpdateTopRequest
}
