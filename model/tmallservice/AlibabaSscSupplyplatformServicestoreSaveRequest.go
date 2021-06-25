package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
保存网点 APIRequest
alibaba.ssc.supplyplatform.servicestore.save

网点创建、修改
*/
type AlibabaSscSupplyplatformServicestoreSaveRequest struct {
    model.Params

    // 入参
    serviceStoreSaveReq   *ServiceStoreSaveForTopReqDTO 

}

func NewAlibabaSscSupplyplatformServicestoreSaveRequest() *AlibabaSscSupplyplatformServicestoreSaveRequest{
    return &AlibabaSscSupplyplatformServicestoreSaveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSscSupplyplatformServicestoreSaveRequest) GetApiMethodName() string {
    return "alibaba.ssc.supplyplatform.servicestore.save"
}

func (r AlibabaSscSupplyplatformServicestoreSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSscSupplyplatformServicestoreSaveRequest) SetServiceStoreSaveReq(serviceStoreSaveReq *ServiceStoreSaveForTopReqDTO) error {
    r.serviceStoreSaveReq = serviceStoreSaveReq
    r.Set("service_store_save_req", serviceStoreSaveReq)
    return nil
}

func (r AlibabaSscSupplyplatformServicestoreSaveRequest) GetServiceStoreSaveReq() *ServiceStoreSaveForTopReqDTO {
    return r.serviceStoreSaveReq
}

