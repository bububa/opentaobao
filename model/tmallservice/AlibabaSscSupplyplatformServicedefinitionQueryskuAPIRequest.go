package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest 服务sku查询 API请求
// alibaba.ssc.supplyplatform.servicedefinition.querysku
//
// 服务sku查询
type AlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest struct {
	model.Params
	// 服务目录id
	_serviceCategoryId string
}

// NewAlibabaSscSupplyplatformServicedefinitionQueryskuRequest 初始化AlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest对象
func NewAlibabaSscSupplyplatformServicedefinitionQueryskuRequest() *AlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest {
	return &AlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest) Reset() {
	r._serviceCategoryId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.servicedefinition.querysku"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceCategoryId is ServiceCategoryId Setter
// 服务目录id
func (r *AlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest) SetServiceCategoryId(_serviceCategoryId string) error {
	r._serviceCategoryId = _serviceCategoryId
	r.Set("service_category_id", _serviceCategoryId)
	return nil
}

// GetServiceCategoryId ServiceCategoryId Getter
func (r AlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest) GetServiceCategoryId() string {
	return r._serviceCategoryId
}

var poolAlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSscSupplyplatformServicedefinitionQueryskuRequest()
	},
}

// GetAlibabaSscSupplyplatformServicedefinitionQueryskuRequest 从 sync.Pool 获取 AlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest
func GetAlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest() *AlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest {
	return poolAlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest.Get().(*AlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest)
}

// ReleaseAlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest 将 AlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest 放入 sync.Pool
func ReleaseAlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest(v *AlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest) {
	v.Reset()
	poolAlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest.Put(v)
}
