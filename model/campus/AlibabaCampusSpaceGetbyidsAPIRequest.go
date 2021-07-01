package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据ids和类型查询空间列表 API请求
alibaba.campus.space.getbyids

根据ids和类型查询空间列表
*/
type AlibabaCampusSpaceGetbyidsAPIRequest struct {
    model.Params
    // 上下文
    _context   *WorkBenchContext
    // 查询条件
    _query   *SpaceIdsQuery
}

// 初始化AlibabaCampusSpaceGetbyidsAPIRequest对象
func NewAlibabaCampusSpaceGetbyidsRequest() *AlibabaCampusSpaceGetbyidsAPIRequest{
    return &AlibabaCampusSpaceGetbyidsAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceGetbyidsAPIRequest) GetApiMethodName() string {
    return "alibaba.campus.space.getbyids"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceGetbyidsAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Context Setter
// 上下文
func (r *AlibabaCampusSpaceGetbyidsAPIRequest) SetContext(_context *WorkBenchContext) error {
    r._context = _context
    r.Set("context", _context)
    return nil
}

// Context Getter
func (r AlibabaCampusSpaceGetbyidsAPIRequest) GetContext() *WorkBenchContext {
    return r._context
}
// Query Setter
// 查询条件
func (r *AlibabaCampusSpaceGetbyidsAPIRequest) SetQuery(_query *SpaceIdsQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r AlibabaCampusSpaceGetbyidsAPIRequest) GetQuery() *SpaceIdsQuery {
    return r._query
}
