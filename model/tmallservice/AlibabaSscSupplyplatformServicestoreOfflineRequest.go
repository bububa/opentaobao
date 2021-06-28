package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
网点下线 APIRequest
alibaba.ssc.supplyplatform.servicestore.offline

网点下线功能
*/
type AlibabaSscSupplyplatformServicestoreOfflineRequest struct {
    model.Params

    // 网点编码列表集合,最大支持1000
    serviceStoreCodeList   []string 

    // 网点id列表集合,最大支持1000
    serviceStoreIdList   []int64 

}

func NewAlibabaSscSupplyplatformServicestoreOfflineRequest() *AlibabaSscSupplyplatformServicestoreOfflineRequest{
    return &AlibabaSscSupplyplatformServicestoreOfflineRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSscSupplyplatformServicestoreOfflineRequest) GetApiMethodName() string {
    return "alibaba.ssc.supplyplatform.servicestore.offline"
}

func (r AlibabaSscSupplyplatformServicestoreOfflineRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSscSupplyplatformServicestoreOfflineRequest) SetServiceStoreCodeList(serviceStoreCodeList []string) error {
    r.serviceStoreCodeList = serviceStoreCodeList
    r.Set("service_store_code_list", serviceStoreCodeList)
    return nil
}

func (r AlibabaSscSupplyplatformServicestoreOfflineRequest) GetServiceStoreCodeList() []string {
    return r.serviceStoreCodeList
}

func (r *AlibabaSscSupplyplatformServicestoreOfflineRequest) SetServiceStoreIdList(serviceStoreIdList []int64) error {
    r.serviceStoreIdList = serviceStoreIdList
    r.Set("service_store_id_list", serviceStoreIdList)
    return nil
}

func (r AlibabaSscSupplyplatformServicestoreOfflineRequest) GetServiceStoreIdList() []int64 {
    return r.serviceStoreIdList
}

