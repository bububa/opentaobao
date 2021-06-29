package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
保存网点 API请求
alibaba.ssc.supplyplatform.servicestore.save

网点创建、修改
*/
type AlibabaSscSupplyplatformServicestoreSaveRequest struct {
    model.Params
    // 入参
    serviceStoreSaveReq   *ServiceStoreSaveForTopReqDTO
}

// 初始化AlibabaSscSupplyplatformServicestoreSaveRequest对象
func NewAlibabaSscSupplyplatformServicestoreSaveRequest() *AlibabaSscSupplyplatformServicestoreSaveRequest{
    return &AlibabaSscSupplyplatformServicestoreSaveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServicestoreSaveRequest) GetApiMethodName() string {
    return "alibaba.ssc.supplyplatform.servicestore.save"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServicestoreSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServiceStoreSaveReq Setter
// 入参
func (r *AlibabaSscSupplyplatformServicestoreSaveRequest) SetServiceStoreSaveReq(serviceStoreSaveReq *ServiceStoreSaveForTopReqDTO) error {
    r.serviceStoreSaveReq = serviceStoreSaveReq
    r.Set("service_store_save_req", serviceStoreSaveReq)
    return nil
}

// ServiceStoreSaveReq Getter
func (r AlibabaSscSupplyplatformServicestoreSaveRequest) GetServiceStoreSaveReq() *ServiceStoreSaveForTopReqDTO {
    return r.serviceStoreSaveReq
}
