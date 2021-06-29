package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
空间单元id查业务属性实例 APIRequest
alibaba.campus.space.unit.getspaceunitwithattr

空间单元id查业务属性实例
*/
type AlibabaCampusSpaceUnitGetspaceunitwithattrRequest struct {
    model.Params

    // 操作用户上下文
    context   *WorkBenchContext 

    // 空间单元id
    spaceUnitId   int64 

}

func NewAlibabaCampusSpaceUnitGetspaceunitwithattrRequest() *AlibabaCampusSpaceUnitGetspaceunitwithattrRequest{
    return &AlibabaCampusSpaceUnitGetspaceunitwithattrRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusSpaceUnitGetspaceunitwithattrRequest) GetApiMethodName() string {
    return "alibaba.campus.space.unit.getspaceunitwithattr"
}

func (r AlibabaCampusSpaceUnitGetspaceunitwithattrRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusSpaceUnitGetspaceunitwithattrRequest) SetContext(context *WorkBenchContext) error {
    r.context = context
    r.Set("context", context)
    return nil
}

func (r AlibabaCampusSpaceUnitGetspaceunitwithattrRequest) GetContext() *WorkBenchContext {
    return r.context
}

func (r *AlibabaCampusSpaceUnitGetspaceunitwithattrRequest) SetSpaceUnitId(spaceUnitId int64) error {
    r.spaceUnitId = spaceUnitId
    r.Set("space_unit_id", spaceUnitId)
    return nil
}

func (r AlibabaCampusSpaceUnitGetspaceunitwithattrRequest) GetSpaceUnitId() int64 {
    return r.spaceUnitId
}

