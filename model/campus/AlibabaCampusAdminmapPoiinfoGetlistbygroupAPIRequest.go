package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIRequest 根据分组条件查询分组下的空间单元不包涵业务属性信息 API请求
// alibaba.campus.adminmap.poiinfo.getlistbygroup
//
// 根据分组条件查询分组下的空间单元不包涵业务属性信息
type AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIRequest struct {
	model.Params
	// 上下文
	_context *WorkBenchContext
	// 查询对象
	_query *SpaceUnitQuery
}

// NewAlibabaCampusAdminmapPoiinfoGetlistbygroupRequest 初始化AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIRequest对象
func NewAlibabaCampusAdminmapPoiinfoGetlistbygroupRequest() *AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIRequest {
	return &AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.adminmap.poiinfo.getlistbygroup"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Context Setter
// 上下文
func (r *AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIRequest) SetContext(_context *WorkBenchContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// Get Context Getter
func (r AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIRequest) GetContext() *WorkBenchContext {
	return r._context
}

// Set is Query Setter
// 查询对象
func (r *AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIRequest) SetQuery(_query *SpaceUnitQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// Get Query Getter
func (r AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIRequest) GetQuery() *SpaceUnitQuery {
	return r._query
}
