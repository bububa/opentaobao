package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterservicestoredeleteservicestorecoverserviceAPIRequest 删除网点覆盖的服务 API请求
// tmall.servicecenter.servicestore.deleteservicestorecoverservice
//
// 天猫服务平台删除网点覆盖的服务，
// 必选字段：serviceStoreCode、bizType
type TmallservicecenterservicestoredeleteservicestorecoverserviceAPIRequest struct {
	model.Params
	// 网点编码
	_serviceStoreCode string
	// 业务类型
	_bizType string
}

// NewTmallservicecenterservicestoredeleteservicestorecoverserviceRequest 初始化TmallservicecenterservicestoredeleteservicestorecoverserviceAPIRequest对象
func NewTmallservicecenterservicestoredeleteservicestorecoverserviceRequest() *TmallservicecenterservicestoredeleteservicestorecoverserviceAPIRequest {
	return &TmallservicecenterservicestoredeleteservicestorecoverserviceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterservicestoredeleteservicestorecoverserviceAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicestore.deleteservicestorecoverservice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterservicestoredeleteservicestorecoverserviceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterservicestoredeleteservicestorecoverserviceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceStoreCode is ServiceStoreCode Setter
// 网点编码
func (r *TmallservicecenterservicestoredeleteservicestorecoverserviceAPIRequest) SetServiceStoreCode(_serviceStoreCode string) error {
	r._serviceStoreCode = _serviceStoreCode
	r.Set("service_store_code", _serviceStoreCode)
	return nil
}

// GetServiceStoreCode ServiceStoreCode Getter
func (r TmallservicecenterservicestoredeleteservicestorecoverserviceAPIRequest) GetServiceStoreCode() string {
	return r._serviceStoreCode
}

// SetBizType is BizType Setter
// 业务类型
func (r *TmallservicecenterservicestoredeleteservicestorecoverserviceAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TmallservicecenterservicestoredeleteservicestorecoverserviceAPIRequest) GetBizType() string {
	return r._bizType
}
