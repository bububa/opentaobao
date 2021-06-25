package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除网点容量 APIRequest
tmall.servicecenter.servicestore.deleteservicestorecapacity

删除网点覆盖的服务，无法恢复，如果只是暂时屏蔽网点的某个能力，可以将此能力对应的网点容量的capacity字段更新为0
必选字段：serviceStoreCode、bizType
*/
type TmallServicecenterServicestoreDeleteservicestorecapacityRequest struct {
    model.Params

    // 网点编码
    serviceStoreCode   string 

    // 业务类型
    bizType   string 

}

func NewTmallServicecenterServicestoreDeleteservicestorecapacityRequest() *TmallServicecenterServicestoreDeleteservicestorecapacityRequest{
    return &TmallServicecenterServicestoreDeleteservicestorecapacityRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterServicestoreDeleteservicestorecapacityRequest) GetApiMethodName() string {
    return "tmall.servicecenter.servicestore.deleteservicestorecapacity"
}

func (r TmallServicecenterServicestoreDeleteservicestorecapacityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterServicestoreDeleteservicestorecapacityRequest) SetServiceStoreCode(serviceStoreCode string) error {
    r.serviceStoreCode = serviceStoreCode
    r.Set("service_store_code", serviceStoreCode)
    return nil
}

func (r TmallServicecenterServicestoreDeleteservicestorecapacityRequest) GetServiceStoreCode() string {
    return r.serviceStoreCode
}

func (r *TmallServicecenterServicestoreDeleteservicestorecapacityRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r TmallServicecenterServicestoreDeleteservicestorecapacityRequest) GetBizType() string {
    return r.bizType
}

