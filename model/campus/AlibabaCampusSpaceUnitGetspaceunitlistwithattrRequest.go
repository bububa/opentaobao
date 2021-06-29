package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
空间单元列表带业务属性实例 APIRequest
alibaba.campus.space.unit.getspaceunitlistwithattr

空间单元列表带业务属性实例
*/
type AlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest struct {
    model.Params

    // 操作用户上下文
    context   *WorkBenchContext 

    // 空间单元查询对象
    unitQuery   *SpaceUnitQuery 

}

func NewAlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest() *AlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest{
    return &AlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest) GetApiMethodName() string {
    return "alibaba.campus.space.unit.getspaceunitlistwithattr"
}

func (r AlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest) SetContext(context *WorkBenchContext) error {
    r.context = context
    r.Set("context", context)
    return nil
}

func (r AlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest) GetContext() *WorkBenchContext {
    return r.context
}

func (r *AlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest) SetUnitQuery(unitQuery *SpaceUnitQuery) error {
    r.unitQuery = unitQuery
    r.Set("unit_query", unitQuery)
    return nil
}

func (r AlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest) GetUnitQuery() *SpaceUnitQuery {
    return r.unitQuery
}

