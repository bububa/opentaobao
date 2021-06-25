package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建门店 APIRequest
tmall.servicecenter.servicestore.create

用于创建门店/网点。多个业务共用
*/
type TmallServicecenterServicestoreCreateRequest struct {
    model.Params

    // 网点/门店
    serviceStore   *ServiceStoreDTO 

}

func NewTmallServicecenterServicestoreCreateRequest() *TmallServicecenterServicestoreCreateRequest{
    return &TmallServicecenterServicestoreCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterServicestoreCreateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.servicestore.create"
}

func (r TmallServicecenterServicestoreCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterServicestoreCreateRequest) SetServiceStore(serviceStore *ServiceStoreDTO) error {
    r.serviceStore = serviceStore
    r.Set("service_store", serviceStore)
    return nil
}

func (r TmallServicecenterServicestoreCreateRequest) GetServiceStore() *ServiceStoreDTO {
    return r.serviceStore
}

