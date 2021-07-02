package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicestoreDeleteservicestorecapacityAPIRequest 删除网点容量 API请求
// tmall.servicecenter.servicestore.deleteservicestorecapacity
//
// 删除网点覆盖的服务，无法恢复，如果只是暂时屏蔽网点的某个能力，可以将此能力对应的网点容量的capacity字段更新为0
// 必选字段：serviceStoreCode、bizType
type TmallServicecenterServicestoreDeleteservicestorecapacityAPIRequest struct {
	model.Params
	// 网点编码
	_serviceStoreCode string
	// 业务类型
	_bizType string
}

// NewTmallServicecenterServicestoreDeleteservicestorecapacityRequest 初始化TmallServicecenterServicestoreDeleteservicestorecapacityAPIRequest对象
func NewTmallServicecenterServicestoreDeleteservicestorecapacityRequest() *TmallServicecenterServicestoreDeleteservicestorecapacityAPIRequest {
	return &TmallServicecenterServicestoreDeleteservicestorecapacityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicestoreDeleteservicestorecapacityAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicestore.deleteservicestorecapacity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicestoreDeleteservicestorecapacityAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetServiceStoreCode is ServiceStoreCode Setter
// 网点编码
func (r *TmallServicecenterServicestoreDeleteservicestorecapacityAPIRequest) SetServiceStoreCode(_serviceStoreCode string) error {
	r._serviceStoreCode = _serviceStoreCode
	r.Set("service_store_code", _serviceStoreCode)
	return nil
}

// GetServiceStoreCode ServiceStoreCode Getter
func (r TmallServicecenterServicestoreDeleteservicestorecapacityAPIRequest) GetServiceStoreCode() string {
	return r._serviceStoreCode
}

// SetBizType is BizType Setter
// 业务类型
func (r *TmallServicecenterServicestoreDeleteservicestorecapacityAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TmallServicecenterServicestoreDeleteservicestorecapacityAPIRequest) GetBizType() string {
	return r._bizType
}
