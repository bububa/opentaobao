package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusadminmappoiinfogetlistbygroupAPIRequest 根据分组条件查询分组下的空间单元不包涵业务属性信息 API请求
// alibaba.campus.adminmap.poiinfo.getlistbygroup
//
// 根据分组条件查询分组下的空间单元不包涵业务属性信息
type AlibabacampusadminmappoiinfogetlistbygroupAPIRequest struct {
	model.Params
	// 上下文
	_context *WorkBenchContext
	// 查询对象
	_query *SpaceUnitQuery
}

// NewAlibabacampusadminmappoiinfogetlistbygroupRequest 初始化AlibabacampusadminmappoiinfogetlistbygroupAPIRequest对象
func NewAlibabacampusadminmappoiinfogetlistbygroupRequest() *AlibabacampusadminmappoiinfogetlistbygroupAPIRequest {
	return &AlibabacampusadminmappoiinfogetlistbygroupAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusadminmappoiinfogetlistbygroupAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.adminmap.poiinfo.getlistbygroup"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusadminmappoiinfogetlistbygroupAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusadminmappoiinfogetlistbygroupAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContext is Context Setter
// 上下文
func (r *AlibabacampusadminmappoiinfogetlistbygroupAPIRequest) SetContext(_context *WorkBenchContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AlibabacampusadminmappoiinfogetlistbygroupAPIRequest) GetContext() *WorkBenchContext {
	return r._context
}

// SetQuery is Query Setter
// 查询对象
func (r *AlibabacampusadminmappoiinfogetlistbygroupAPIRequest) SetQuery(_query *SpaceUnitQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabacampusadminmappoiinfogetlistbygroupAPIRequest) GetQuery() *SpaceUnitQuery {
	return r._query
}
