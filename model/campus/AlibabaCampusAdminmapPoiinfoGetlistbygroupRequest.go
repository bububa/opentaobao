package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据分组条件查询分组下的空间单元不包涵业务属性信息 API请求
alibaba.campus.adminmap.poiinfo.getlistbygroup

根据分组条件查询分组下的空间单元不包涵业务属性信息
*/
type AlibabaCampusAdminmapPoiinfoGetlistbygroupRequest struct {
    model.Params
    // 上下文
    _context   *WorkBenchContext
    // 查询对象
    _query   *SpaceUnitQuery
}

// 初始化AlibabaCampusAdminmapPoiinfoGetlistbygroupRequest对象
func NewAlibabaCampusAdminmapPoiinfoGetlistbygroupRequest() *AlibabaCampusAdminmapPoiinfoGetlistbygroupRequest{
    return &AlibabaCampusAdminmapPoiinfoGetlistbygroupRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAdminmapPoiinfoGetlistbygroupRequest) GetApiMethodName() string {
    return "alibaba.campus.adminmap.poiinfo.getlistbygroup"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAdminmapPoiinfoGetlistbygroupRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Context Setter
// 上下文
func (r *AlibabaCampusAdminmapPoiinfoGetlistbygroupRequest) SetContext(_context *WorkBenchContext) error {
    r._context = _context
    r.Set("context", _context)
    return nil
}

// Context Getter
func (r AlibabaCampusAdminmapPoiinfoGetlistbygroupRequest) GetContext() *WorkBenchContext {
    return r._context
}
// Query Setter
// 查询对象
func (r *AlibabaCampusAdminmapPoiinfoGetlistbygroupRequest) SetQuery(_query *SpaceUnitQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r AlibabaCampusAdminmapPoiinfoGetlistbygroupRequest) GetQuery() *SpaceUnitQuery {
    return r._query
}
