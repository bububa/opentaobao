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
type AlibabaCampusSpaceGetbyidsRequest struct {
    model.Params
    // 上下文
    context   *WorkBenchContext
    // 查询条件
    query   *SpaceIdsQuery
}

// 初始化AlibabaCampusSpaceGetbyidsRequest对象
func NewAlibabaCampusSpaceGetbyidsRequest() *AlibabaCampusSpaceGetbyidsRequest{
    return &AlibabaCampusSpaceGetbyidsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceGetbyidsRequest) GetApiMethodName() string {
    return "alibaba.campus.space.getbyids"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceGetbyidsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Context Setter
// 上下文
func (r *AlibabaCampusSpaceGetbyidsRequest) SetContext(context *WorkBenchContext) error {
    r.context = context
    r.Set("context", context)
    return nil
}

// Context Getter
func (r AlibabaCampusSpaceGetbyidsRequest) GetContext() *WorkBenchContext {
    return r.context
}
// Query Setter
// 查询条件
func (r *AlibabaCampusSpaceGetbyidsRequest) SetQuery(query *SpaceIdsQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r AlibabaCampusSpaceGetbyidsRequest) GetQuery() *SpaceIdsQuery {
    return r.query
}
