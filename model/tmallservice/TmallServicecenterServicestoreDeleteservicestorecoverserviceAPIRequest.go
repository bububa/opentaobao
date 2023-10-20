package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest 删除网点覆盖的服务 API请求
// tmall.servicecenter.servicestore.deleteservicestorecoverservice
//
// 天猫服务平台删除网点覆盖的服务，
// 必选字段：serviceStoreCode、bizType
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest) Reset() {
	r._serviceStoreCode = ""
	r._bizType = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicestore.deleteservicestorecoverservice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceStoreCode is ServiceStoreCode Setter
// 网点编码
func (r *TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest) SetServiceStoreCode(_serviceStoreCode string) error {
	r._serviceStoreCode = _serviceStoreCode
	r.Set("service_store_code", _serviceStoreCode)
	return nil
}

// GetServiceStoreCode ServiceStoreCode Getter
func (r TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest) GetServiceStoreCode() string {
	return r._serviceStoreCode
}

// SetBizType is BizType Setter
// 业务类型
func (r *TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest) GetBizType() string {
	return r._bizType
}

var poolTmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterServicestoreDeleteservicestorecoverserviceRequest()
	},
}

// GetTmallServicecenterServicestoreDeleteservicestorecoverserviceRequest 从 sync.Pool 获取 TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest
func GetTmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest() *TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest {
	return poolTmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest.Get().(*TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest)
}

// ReleaseTmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest 将 TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest(v *TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest) {
	v.Reset()
	poolTmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest.Put(v)
}
