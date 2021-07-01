package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest
删除网点覆盖的服务 API请求
tmall.servicecenter.servicestore.deleteservicestorecoverservice

天猫服务平台删除网点覆盖的服务，
必选字段：serviceStoreCode、bizType */
type TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest struct {
	model.Params
	// 网点编码
	_serviceStoreCode string
	// 业务类型
	_bizType string
}

// NewTmallServicecenterServicestoreDeleteservicestorecoverserviceRequest 初始化TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest对象
func NewTmallServicecenterServicestoreDeleteservicestorecoverserviceRequest() *TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest {
	return &TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicestore.deleteservicestorecoverservice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ServiceStoreCode Setter
// 网点编码
func (r *TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest) SetServiceStoreCode(_serviceStoreCode string) error {
	r._serviceStoreCode = _serviceStoreCode
	r.Set("service_store_code", _serviceStoreCode)
	return nil
}

// Get ServiceStoreCode Getter
func (r TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest) GetServiceStoreCode() string {
	return r._serviceStoreCode
}

// Set is BizType Setter
// 业务类型
func (r *TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// Get BizType Getter
func (r TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest) GetBizType() string {
	return r._bizType
}
