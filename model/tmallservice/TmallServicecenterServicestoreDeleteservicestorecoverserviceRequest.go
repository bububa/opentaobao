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
    serviceStoreCode   string
    // 业务类型
    bizType   string
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
func (r *TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest) SetServiceStoreCode(serviceStoreCode string) error {
    r.serviceStoreCode = serviceStoreCode
    r.Set("service_store_code", serviceStoreCode)
    return nil
}

// ServiceStoreCode Getter
func (r TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest) GetServiceStoreCode() string {
    return r.serviceStoreCode
}
// BizType Setter
// 业务类型
func (r *TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest) GetBizType() string {
    return r.bizType
}
