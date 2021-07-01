package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据id获取楼层 API请求
alibaba.campus.space.floor.getbyid

根据id获取楼层
*/
type AlibabaCampusSpaceFloorGetbyidAPIRequest struct {
    model.Params
    // 环境上下文
    _context   *WorkBenchContext
    // 楼层id
    _id   int64
}

// 初始化AlibabaCampusSpaceFloorGetbyidAPIRequest对象
func NewAlibabaCampusSpaceFloorGetbyidRequest() *AlibabaCampusSpaceFloorGetbyidAPIRequest{
    return &AlibabaCampusSpaceFloorGetbyidAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceFloorGetbyidAPIRequest) GetApiMethodName() string {
    return "alibaba.campus.space.floor.getbyid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceFloorGetbyidAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Context Setter
// 环境上下文
func (r *AlibabaCampusSpaceFloorGetbyidAPIRequest) SetContext(_context *WorkBenchContext) error {
    r._context = _context
    r.Set("context", _context)
    return nil
}

// Context Getter
func (r AlibabaCampusSpaceFloorGetbyidAPIRequest) GetContext() *WorkBenchContext {
    return r._context
}
// Id Setter
// 楼层id
func (r *AlibabaCampusSpaceFloorGetbyidAPIRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaCampusSpaceFloorGetbyidAPIRequest) GetId() int64 {
    return r._id
}
