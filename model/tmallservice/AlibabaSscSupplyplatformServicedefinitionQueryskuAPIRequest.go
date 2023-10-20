package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscsupplyplatformservicedefinitionqueryskuAPIRequest 服务sku查询 API请求
// alibaba.ssc.supplyplatform.servicedefinition.querysku
//
// 服务sku查询
type AlibabasscsupplyplatformservicedefinitionqueryskuAPIRequest struct {
	model.Params
	// 服务目录id
	_serviceCategoryId string
}

// NewAlibabasscsupplyplatformservicedefinitionqueryskuRequest 初始化AlibabasscsupplyplatformservicedefinitionqueryskuAPIRequest对象
func NewAlibabasscsupplyplatformservicedefinitionqueryskuRequest() *AlibabasscsupplyplatformservicedefinitionqueryskuAPIRequest {
	return &AlibabasscsupplyplatformservicedefinitionqueryskuAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasscsupplyplatformservicedefinitionqueryskuAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.servicedefinition.querysku"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasscsupplyplatformservicedefinitionqueryskuAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasscsupplyplatformservicedefinitionqueryskuAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceCategoryId is ServiceCategoryId Setter
// 服务目录id
func (r *AlibabasscsupplyplatformservicedefinitionqueryskuAPIRequest) SetServiceCategoryId(_serviceCategoryId string) error {
	r._serviceCategoryId = _serviceCategoryId
	r.Set("service_category_id", _serviceCategoryId)
	return nil
}

// GetServiceCategoryId ServiceCategoryId Getter
func (r AlibabasscsupplyplatformservicedefinitionqueryskuAPIRequest) GetServiceCategoryId() string {
	return r._serviceCategoryId
}
