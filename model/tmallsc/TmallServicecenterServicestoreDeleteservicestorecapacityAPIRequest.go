package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterservicestoredeleteservicestorecapacityAPIRequest 删除网点容量 API请求
// tmall.servicecenter.servicestore.deleteservicestorecapacity
//
// 删除网点覆盖的服务，无法恢复，如果只是暂时屏蔽网点的某个能力，可以将此能力对应的网点容量的capacity字段更新为0
// 必选字段：serviceStoreCode、bizType
type TmallservicecenterservicestoredeleteservicestorecapacityAPIRequest struct {
	model.Params
	// 网点编码
	_serviceStoreCode string
	// 业务类型
	_bizType string
}

// NewTmallservicecenterservicestoredeleteservicestorecapacityRequest 初始化TmallservicecenterservicestoredeleteservicestorecapacityAPIRequest对象
func NewTmallservicecenterservicestoredeleteservicestorecapacityRequest() *TmallservicecenterservicestoredeleteservicestorecapacityAPIRequest {
	return &TmallservicecenterservicestoredeleteservicestorecapacityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterservicestoredeleteservicestorecapacityAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicestore.deleteservicestorecapacity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterservicestoredeleteservicestorecapacityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterservicestoredeleteservicestorecapacityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceStoreCode is ServiceStoreCode Setter
// 网点编码
func (r *TmallservicecenterservicestoredeleteservicestorecapacityAPIRequest) SetServiceStoreCode(_serviceStoreCode string) error {
	r._serviceStoreCode = _serviceStoreCode
	r.Set("service_store_code", _serviceStoreCode)
	return nil
}

// GetServiceStoreCode ServiceStoreCode Getter
func (r TmallservicecenterservicestoredeleteservicestorecapacityAPIRequest) GetServiceStoreCode() string {
	return r._serviceStoreCode
}

// SetBizType is BizType Setter
// 业务类型
func (r *TmallservicecenterservicestoredeleteservicestorecapacityAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TmallservicecenterservicestoredeleteservicestorecapacityAPIRequest) GetBizType() string {
	return r._bizType
}
