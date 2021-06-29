package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
空间单元id查业务属性实例 API请求
alibaba.campus.space.unit.getspaceunitwithattr

空间单元id查业务属性实例
*/
type AlibabaCampusSpaceUnitGetspaceunitwithattrRequest struct {
    model.Params
    // 操作用户上下文
    _context   *WorkBenchContext
    // 空间单元id
    _spaceUnitId   int64
}

// 初始化AlibabaCampusSpaceUnitGetspaceunitwithattrRequest对象
func NewAlibabaCampusSpaceUnitGetspaceunitwithattrRequest() *AlibabaCampusSpaceUnitGetspaceunitwithattrRequest{
    return &AlibabaCampusSpaceUnitGetspaceunitwithattrRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceUnitGetspaceunitwithattrRequest) GetApiMethodName() string {
    return "alibaba.campus.space.unit.getspaceunitwithattr"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceUnitGetspaceunitwithattrRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Context Setter
// 操作用户上下文
func (r *AlibabaCampusSpaceUnitGetspaceunitwithattrRequest) SetContext(_context *WorkBenchContext) error {
    r._context = _context
    r.Set("context", _context)
    return nil
}

// Context Getter
func (r AlibabaCampusSpaceUnitGetspaceunitwithattrRequest) GetContext() *WorkBenchContext {
    return r._context
}
// SpaceUnitId Setter
// 空间单元id
func (r *AlibabaCampusSpaceUnitGetspaceunitwithattrRequest) SetSpaceUnitId(_spaceUnitId int64) error {
    r._spaceUnitId = _spaceUnitId
    r.Set("space_unit_id", _spaceUnitId)
    return nil
}

// SpaceUnitId Getter
func (r AlibabaCampusSpaceUnitGetspaceunitwithattrRequest) GetSpaceUnitId() int64 {
    return r._spaceUnitId
}
