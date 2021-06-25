package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除网点覆盖的服务 APIRequest
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

func NewTmallServicecenterServicestoreDeleteservicestorecoverserviceRequest() *TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest{
    return &TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest) GetApiMethodName() string {
    return "tmall.servicecenter.servicestore.deleteservicestorecoverservice"
}

func (r TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest) SetServiceStoreCode(serviceStoreCode string) error {
    r.serviceStoreCode = serviceStoreCode
    r.Set("service_store_code", serviceStoreCode)
    return nil
}

func (r TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest) GetServiceStoreCode() string {
    return r.serviceStoreCode
}

func (r *TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r TmallServicecenterServicestoreDeleteservicestorecoverserviceRequest) GetBizType() string {
    return r.bizType
}

