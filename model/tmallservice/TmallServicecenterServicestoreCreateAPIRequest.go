package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicestoreCreateAPIRequest 创建门店 API请求
// tmall.servicecenter.servicestore.create
//
// 用于创建门店/网点。多个业务共用
type TmallServicecenterServicestoreCreateAPIRequest struct {
	model.Params
	// 网点/门店
	_serviceStore *ServiceStoreDto
}

// NewTmallServicecenterServicestoreCreateRequest 初始化TmallServicecenterServicestoreCreateAPIRequest对象
func NewTmallServicecenterServicestoreCreateRequest() *TmallServicecenterServicestoreCreateAPIRequest {
	return &TmallServicecenterServicestoreCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicestoreCreateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicestore.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicestoreCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetServiceStore is ServiceStore Setter
// 网点/门店
func (r *TmallServicecenterServicestoreCreateAPIRequest) SetServiceStore(_serviceStore *ServiceStoreDto) error {
	r._serviceStore = _serviceStore
	r.Set("service_store", _serviceStore)
	return nil
}

// GetServiceStore ServiceStore Getter
func (r TmallServicecenterServicestoreCreateAPIRequest) GetServiceStore() *ServiceStoreDto {
	return r._serviceStore
}
