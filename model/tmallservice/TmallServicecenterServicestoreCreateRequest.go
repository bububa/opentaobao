package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建门店 API请求
tmall.servicecenter.servicestore.create

用于创建门店/网点。多个业务共用
*/
type TmallServicecenterServicestoreCreateRequest struct {
    model.Params
    // 网点/门店
    serviceStore   *ServiceStoreDTO
}

// 初始化TmallServicecenterServicestoreCreateRequest对象
func NewTmallServicecenterServicestoreCreateRequest() *TmallServicecenterServicestoreCreateRequest{
    return &TmallServicecenterServicestoreCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicestoreCreateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.servicestore.create"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicestoreCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServiceStore Setter
// 网点/门店
func (r *TmallServicecenterServicestoreCreateRequest) SetServiceStore(serviceStore *ServiceStoreDTO) error {
    r.serviceStore = serviceStore
    r.Set("service_store", serviceStore)
    return nil
}

// ServiceStore Getter
func (r TmallServicecenterServicestoreCreateRequest) GetServiceStore() *ServiceStoreDTO {
    return r.serviceStore
}
