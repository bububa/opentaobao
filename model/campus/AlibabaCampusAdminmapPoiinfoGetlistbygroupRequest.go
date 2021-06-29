package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据分组条件查询分组下的空间单元不包涵业务属性信息 APIRequest
alibaba.campus.adminmap.poiinfo.getlistbygroup

根据分组条件查询分组下的空间单元不包涵业务属性信息
*/
type AlibabaCampusAdminmapPoiinfoGetlistbygroupRequest struct {
    model.Params

    // 上下文
    context   *WorkBenchContext 

    // 查询对象
    query   *SpaceUnitQuery 

}

func NewAlibabaCampusAdminmapPoiinfoGetlistbygroupRequest() *AlibabaCampusAdminmapPoiinfoGetlistbygroupRequest{
    return &AlibabaCampusAdminmapPoiinfoGetlistbygroupRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAdminmapPoiinfoGetlistbygroupRequest) GetApiMethodName() string {
    return "alibaba.campus.adminmap.poiinfo.getlistbygroup"
}

func (r AlibabaCampusAdminmapPoiinfoGetlistbygroupRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAdminmapPoiinfoGetlistbygroupRequest) SetContext(context *WorkBenchContext) error {
    r.context = context
    r.Set("context", context)
    return nil
}

func (r AlibabaCampusAdminmapPoiinfoGetlistbygroupRequest) GetContext() *WorkBenchContext {
    return r.context
}

func (r *AlibabaCampusAdminmapPoiinfoGetlistbygroupRequest) SetQuery(query *SpaceUnitQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r AlibabaCampusAdminmapPoiinfoGetlistbygroupRequest) GetQuery() *SpaceUnitQuery {
    return r.query
}

