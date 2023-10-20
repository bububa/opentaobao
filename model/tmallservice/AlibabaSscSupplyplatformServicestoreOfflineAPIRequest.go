package tmallservice

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSscSupplyplatformServicestoreOfflineAPIRequest) Reset() {
	r._serviceStoreCodeList = r._serviceStoreCodeList[:0]
	r._serviceStoreIdList = r._serviceStoreIdList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServicestoreOfflineAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.servicestore.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServicestoreOfflineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSscSupplyplatformServicestoreOfflineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceStoreCodeList is ServiceStoreCodeList Setter
// 网点编码列表集合,最大支持1000
func (r *AlibabaSscSupplyplatformServicestoreOfflineAPIRequest) SetServiceStoreCodeList(_serviceStoreCodeList []string) error {
	r._serviceStoreCodeList = _serviceStoreCodeList
	r.Set("service_store_code_list", _serviceStoreCodeList)
	return nil
}

// GetServiceStoreCodeList ServiceStoreCodeList Getter
func (r AlibabaSscSupplyplatformServicestoreOfflineAPIRequest) GetServiceStoreCodeList() []string {
	return r._serviceStoreCodeList
}

// SetServiceStoreIdList is ServiceStoreIdList Setter
// 网点id列表集合,最大支持1000
func (r *AlibabaSscSupplyplatformServicestoreOfflineAPIRequest) SetServiceStoreIdList(_serviceStoreIdList []int64) error {
	r._serviceStoreIdList = _serviceStoreIdList
	r.Set("service_store_id_list", _serviceStoreIdList)
	return nil
}

// GetServiceStoreIdList ServiceStoreIdList Getter
func (r AlibabaSscSupplyplatformServicestoreOfflineAPIRequest) GetServiceStoreIdList() []int64 {
	return r._serviceStoreIdList
}

var poolAlibabaSscSupplyplatformServicestoreOfflineAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSscSupplyplatformServicestoreOfflineRequest()
	},
}

// GetAlibabaSscSupplyplatformServicestoreOfflineRequest 从 sync.Pool 获取 AlibabaSscSupplyplatformServicestoreOfflineAPIRequest
func GetAlibabaSscSupplyplatformServicestoreOfflineAPIRequest() *AlibabaSscSupplyplatformServicestoreOfflineAPIRequest {
	return poolAlibabaSscSupplyplatformServicestoreOfflineAPIRequest.Get().(*AlibabaSscSupplyplatformServicestoreOfflineAPIRequest)
}

// ReleaseAlibabaSscSupplyplatformServicestoreOfflineAPIRequest 将 AlibabaSscSupplyplatformServicestoreOfflineAPIRequest 放入 sync.Pool
func ReleaseAlibabaSscSupplyplatformServicestoreOfflineAPIRequest(v *AlibabaSscSupplyplatformServicestoreOfflineAPIRequest) {
	v.Reset()
	poolAlibabaSscSupplyplatformServicestoreOfflineAPIRequest.Put(v)
}
