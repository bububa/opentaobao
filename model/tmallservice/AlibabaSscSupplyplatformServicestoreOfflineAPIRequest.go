package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscsupplyplatformservicestoreofflineAPIRequest 网点下线 API请求
// alibaba.ssc.supplyplatform.servicestore.offline
//
// 网点下线功能
type AlibabasscsupplyplatformservicestoreofflineAPIRequest struct {
	model.Params
	// 网点编码列表集合,最大支持1000
	_serviceStoreCodeList []string
	// 网点id列表集合,最大支持1000
	_serviceStoreIdList []int64
}

// NewAlibabasscsupplyplatformservicestoreofflineRequest 初始化AlibabasscsupplyplatformservicestoreofflineAPIRequest对象
func NewAlibabasscsupplyplatformservicestoreofflineRequest() *AlibabasscsupplyplatformservicestoreofflineAPIRequest {
	return &AlibabasscsupplyplatformservicestoreofflineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasscsupplyplatformservicestoreofflineAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.servicestore.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasscsupplyplatformservicestoreofflineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasscsupplyplatformservicestoreofflineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceStoreCodeList is ServiceStoreCodeList Setter
// 网点编码列表集合,最大支持1000
func (r *AlibabasscsupplyplatformservicestoreofflineAPIRequest) SetServiceStoreCodeList(_serviceStoreCodeList []string) error {
	r._serviceStoreCodeList = _serviceStoreCodeList
	r.Set("service_store_code_list", _serviceStoreCodeList)
	return nil
}

// GetServiceStoreCodeList ServiceStoreCodeList Getter
func (r AlibabasscsupplyplatformservicestoreofflineAPIRequest) GetServiceStoreCodeList() []string {
	return r._serviceStoreCodeList
}

// SetServiceStoreIdList is ServiceStoreIdList Setter
// 网点id列表集合,最大支持1000
func (r *AlibabasscsupplyplatformservicestoreofflineAPIRequest) SetServiceStoreIdList(_serviceStoreIdList []int64) error {
	r._serviceStoreIdList = _serviceStoreIdList
	r.Set("service_store_id_list", _serviceStoreIdList)
	return nil
}

// GetServiceStoreIdList ServiceStoreIdList Getter
func (r AlibabasscsupplyplatformservicestoreofflineAPIRequest) GetServiceStoreIdList() []int64 {
	return r._serviceStoreIdList
}
