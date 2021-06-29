package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除网点覆盖的服务 API请求
tmall.servicecenter.servicestore.deleteservicestorecoverservice

天猫服务平台删除网点覆盖的服务，
必选字段：serviceStoreCode、bizType
*/
type TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest struct {
    model.Params
    // 网点编码
    _serviceStoreCode   string
    // 业务类型
    _bizType   string
}

// 初始化TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest对象
func NewTmallServicecenterServicestoreDeleteservicestorecoverserviceRequest() *TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest{
    return &TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest) GetApiMethodName() string {
    return "tmall.servicecenter.servicestore.deleteservicestorecoverservice"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServiceStoreCode Setter
// 网点编码
func (r *TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest) SetServiceStoreCode(_serviceStoreCode string) error {
    r._serviceStoreCode = _serviceStoreCode
    r.Set("service_store_code", _serviceStoreCode)
    return nil
}

// ServiceStoreCode Getter
func (r TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest) GetServiceStoreCode() string {
    return r._serviceStoreCode
}
// BizType Setter
// 业务类型
func (r *TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest) SetBizType(_bizType string) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest) GetBizType() string {
    return r._bizType
}
