package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据空间分组id、ids查空间单元信息【带空间单元业务属性信息】 API请求
alibaba.campus.space.unit.getlistwithattrbygroupid

根据空间分组id、ids查空间单元信息【带空间单元业务属性信息】
*/
type AlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest struct {
    model.Params
    // 操作用户上下文
    _context   *WorkBenchContext
    // 分组id
    _groupId   int64
}

// 初始化AlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest对象
func NewAlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest() *AlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest{
    return &AlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest) GetApiMethodName() string {
    return "alibaba.campus.space.unit.getlistwithattrbygroupid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Context Setter
// 操作用户上下文
func (r *AlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest) SetContext(_context *WorkBenchContext) error {
    r._context = _context
    r.Set("context", _context)
    return nil
}

// Context Getter
func (r AlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest) GetContext() *WorkBenchContext {
    return r._context
}
// GroupId Setter
// 分组id
func (r *AlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest) SetGroupId(_groupId int64) error {
    r._groupId = _groupId
    r.Set("group_id", _groupId)
    return nil
}

// GroupId Getter
func (r AlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest) GetGroupId() int64 {
    return r._groupId
}