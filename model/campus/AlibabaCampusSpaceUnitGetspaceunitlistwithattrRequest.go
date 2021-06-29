package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
空间单元列表带业务属性实例 API请求
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

// 初始化AlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest对象
func NewAlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest() *AlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest{
    return &AlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest) GetApiMethodName() string {
    return "alibaba.campus.space.unit.getspaceunitlistwithattr"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Context Setter
// 操作用户上下文
func (r *AlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest) SetContext(context *WorkBenchContext) error {
    r.context = context
    r.Set("context", context)
    return nil
}

// Context Getter
func (r AlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest) GetContext() *WorkBenchContext {
    return r.context
}
// UnitQuery Setter
// 空间单元查询对象
func (r *AlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest) SetUnitQuery(unitQuery *SpaceUnitQuery) error {
    r.unitQuery = unitQuery
    r.Set("unit_query", unitQuery)
    return nil
}

// UnitQuery Getter
func (r AlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest) GetUnitQuery() *SpaceUnitQuery {
    return r.unitQuery
}
