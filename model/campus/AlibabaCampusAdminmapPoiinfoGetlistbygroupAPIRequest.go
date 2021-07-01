package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIRequest
根据分组条件查询分组下的空间单元不包涵业务属性信息 API请求
alibaba.campus.adminmap.poiinfo.getlistbygroup

根据分组条件查询分组下的空间单元不包涵业务属性信息 */
type AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIRequest struct {
	model.Params
	// 上下文
	_context *WorkBenchContext
	// 查询对象
	_query *SpaceUnitQuery
}

// New
