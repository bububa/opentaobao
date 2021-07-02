package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServicestoreOfflineAPIRequest 网点下线 API请求
// alibaba.ssc.supplyplatform.servicestore.offline
//
// 网点下线功能
type AlibabaSscSupplyplatformServicestoreOfflineAPIRequest struct {
	model.Params
	// 网点编码列表集合,最大支持1000
	_serviceStoreCodeList []string
	// 网点id列表集合,最大支持1000
	_serviceStoreIdList []int64
}

// NewAlibabaSscSupplyplatformServicestoreOfflineRequest 初始化AlibabaSscSupplyplatformServicestoreOfflineAPIRequest对象
func NewAlibabaSscSupplyplatformServicestoreOfflineRequest() *AlibabaSscSupplyplatformServicestoreOfflineAPIRequest {
	return &AlibabaSscSupplyplatformServicestoreOfflineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServicestoreOfflineAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.servicestore.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServicestoreOfflineAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ServiceStoreCodeList Setter
// 网点编码列表集合,最大支持1000
func (r *AlibabaSscSupplyplatformServicestoreOfflineAPIRequest) SetServiceStoreCodeList(_serviceStoreCodeList []string) error {
	r._serviceStoreCodeList = _serviceStoreCodeList
	r.Set("service_store_code_list", _serviceStoreCodeList)
	return nil
}

// Get ServiceStoreCodeList Getter
func (r AlibabaSscSupplyplatformServicestoreOfflineAPIRequest) GetServiceStoreCodeList() []string {
	return r._serviceStoreCodeList
}

// Set is ServiceStoreIdList Setter
// 网点id列表集合,最大支持1000
func (r *AlibabaSscSupplyplatformServicestoreOfflineAPIRequest) SetServiceStoreIdList(_serviceStoreIdList []int64) error {
	r._serviceStoreIdList = _serviceStoreIdList
	r.Set("service_store_id_list", _serviceStoreIdList)
	return nil
}

// Get ServiceStoreIdList Getter
func (r AlibabaSscSupplyplatformServicestoreOfflineAPIRequest) GetServiceStoreIdList() []int64 {
	return r._serviceStoreIdList
}
