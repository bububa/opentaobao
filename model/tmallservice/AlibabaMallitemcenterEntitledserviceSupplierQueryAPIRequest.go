package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest
根据天猫id查询门店服务授权 API请求
alibaba.mallitemcenter.entitledservice.supplier.query

根据天猫id查询门店服务授权 */
type AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest struct {
	model.Params
	// 天猫id
	_id int64
	// 第几页
	_currentPage int64
	// 每页条数
	_pageSize int64
}

// New
