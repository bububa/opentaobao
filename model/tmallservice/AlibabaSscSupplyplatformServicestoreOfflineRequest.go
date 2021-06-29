package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
网点下线 API请求
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

// 初始化AlibabaSscSupplyplatformServicestoreOfflineRequest对象
func NewAlibabaSscSupplyplatformServicestoreOfflineRequest() *AlibabaSscSupplyplatformServicestoreOfflineRequest{
    return &AlibabaSscSupplyplatformServicestoreOfflineRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServicestoreOfflineRequest) GetApiMethodName() string {
    return "alibaba.ssc.supplyplatform.servicestore.offline"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServicestoreOfflineRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServiceStoreCodeList Setter
// 网点编码列表集合,最大支持1000
func (r *AlibabaSscSupplyplatformServicestoreOfflineRequest) SetServiceStoreCodeList(serviceStoreCodeList []string) error {
    r.serviceStoreCodeList = serviceStoreCodeList
    r.Set("service_store_code_list", serviceStoreCodeList)
    return nil
}

// ServiceStoreCodeList Getter
func (r AlibabaSscSupplyplatformServicestoreOfflineRequest) GetServiceStoreCodeList() []string {
    return r.serviceStoreCodeList
}
// ServiceStoreIdList Setter
// 网点id列表集合,最大支持1000
func (r *AlibabaSscSupplyplatformServicestoreOfflineRequest) SetServiceStoreIdList(serviceStoreIdList []int64) error {
    r.serviceStoreIdList = serviceStoreIdList
    r.Set("service_store_id_list", serviceStoreIdList)
    return nil
}

// ServiceStoreIdList Getter
func (r AlibabaSscSupplyplatformServicestoreOfflineRequest) GetServiceStoreIdList() []int64 {
    return r.serviceStoreIdList
}
