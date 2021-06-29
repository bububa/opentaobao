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
type AlibabaCampusSpaceFloorGetbyidRequest struct {
    model.Params
    // 环境上下文
    context   *WorkBenchContext
    // 楼层id
    id   int64
}

// 初始化AlibabaCampusSpaceFloorGetbyidRequest对象
func NewAlibabaCampusSpaceFloorGetbyidRequest() *AlibabaCampusSpaceFloorGetbyidRequest{
    return &AlibabaCampusSpaceFloorGetbyidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceFloorGetbyidRequest) GetApiMethodName() string {
    return "alibaba.campus.space.floor.getbyid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceFloorGetbyidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Context Setter
// 环境上下文
func (r *AlibabaCampusSpaceFloorGetbyidRequest) SetContext(context *WorkBenchContext) error {
    r.context = context
    r.Set("context", context)
    return nil
}

// Context Getter
func (r AlibabaCampusSpaceFloorGetbyidRequest) GetContext() *WorkBenchContext {
    return r.context
}
// Id Setter
// 楼层id
func (r *AlibabaCampusSpaceFloorGetbyidRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r AlibabaCampusSpaceFloorGetbyidRequest) GetId() int64 {
    return r.id
}
