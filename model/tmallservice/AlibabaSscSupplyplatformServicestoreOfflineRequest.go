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
    _serviceStoreCodeList   []string
    // 网点id列表集合,最大支持1000
    _serviceStoreIdList   []int64
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
func (r *AlibabaSscSupplyplatformServicestoreOfflineRequest) SetServiceStoreCodeList(_serviceStoreCodeList []string) error {
    r._serviceStoreCodeList = _serviceStoreCodeList
    r.Set("service_store_code_list", _serviceStoreCodeList)
    return nil
}

// ServiceStoreCodeList Getter
func (r AlibabaSscSupplyplatformServicestoreOfflineRequest) GetServiceStoreCodeList() []string {
    return r._serviceStoreCodeList
}
// ServiceStoreIdList Setter
// 网点id列表集合,最大支持1000
func (r *AlibabaSscSupplyplatformServicestoreOfflineRequest) SetServiceStoreIdList(_serviceStoreIdList []int64) error {
    r._serviceStoreIdList = _serviceStoreIdList
    r.Set("service_store_id_list", _serviceStoreIdList)
    return nil
}

// ServiceStoreIdList Getter
func (r AlibabaSscSupplyplatformServicestoreOfflineRequest) GetServiceStoreIdList() []int64 {
    return r._serviceStoreIdList
}
