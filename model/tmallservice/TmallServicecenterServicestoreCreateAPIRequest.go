package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterservicestorecreateAPIRequest 创建门店 API请求
// tmall.servicecenter.servicestore.create
//
// 用于创建门店/网点。多个业务共用
type TmallservicecenterservicestorecreateAPIRequest struct {
	model.Params
	// 网点/门店
	_serviceStore *ServiceStoreDto
}

// NewTmallservicecenterservicestorecreateRequest 初始化TmallservicecenterservicestorecreateAPIRequest对象
func NewTmallservicecenterservicestorecreateRequest() *TmallservicecenterservicestorecreateAPIRequest {
	return &TmallservicecenterservicestorecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterservicestorecreateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicestore.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterservicestorecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterservicestorecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceStore is ServiceStore Setter
// 网点/门店
func (r *TmallservicecenterservicestorecreateAPIRequest) SetServiceStore(_serviceStore *ServiceStoreDto) error {
	r._serviceStore = _serviceStore
	r.Set("service_store", _serviceStore)
	return nil
}

// GetServiceStore ServiceStore Getter
func (r TmallservicecenterservicestorecreateAPIRequest) GetServiceStore() *ServiceStoreDto {
	return r._serviceStore
}
