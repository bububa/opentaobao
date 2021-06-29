package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据ids和类型查询空间列表 APIRequest
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

func NewAlibabaCampusSpaceGetbyidsRequest() *AlibabaCampusSpaceGetbyidsRequest{
    return &AlibabaCampusSpaceGetbyidsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusSpaceGetbyidsRequest) GetApiMethodName() string {
    return "alibaba.campus.space.getbyids"
}

func (r AlibabaCampusSpaceGetbyidsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusSpaceGetbyidsRequest) SetContext(context *WorkBenchContext) error {
    r.context = context
    r.Set("context", context)
    return nil
}

func (r AlibabaCampusSpaceGetbyidsRequest) GetContext() *WorkBenchContext {
    return r.context
}

func (r *AlibabaCampusSpaceGetbyidsRequest) SetQuery(query *SpaceIdsQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r AlibabaCampusSpaceGetbyidsRequest) GetQuery() *SpaceIdsQuery {
    return r.query
}

