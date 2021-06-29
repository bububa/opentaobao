package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询活动列表 API请求
taobao.ump.activities.get

查询活动列表
*/
type TaobaoUmpActivitiesGetRequest struct {
    model.Params
    // 工具id
    toolId   int64
    // 分页的页码
    pageNo   int64
    // 每页的最大条数
    pageSize   int64
}

// 初始化TaobaoUmpActivitiesGetRequest对象
func NewTaobaoUmpActivitiesGetRequest() *TaobaoUmpActivitiesGetRequest{
    return &TaobaoUmpActivitiesGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpActivitiesGetRequest) GetApiMethodName() string {
    return "taobao.ump.activities.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpActivitiesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ToolId Setter
// 工具id
func (r *TaobaoUmpActivitiesGetRequest) SetToolId(toolId int64) error {
    r.toolId = toolId
    r.Set("tool_id", toolId)
    return nil
}

// ToolId Getter
func (r TaobaoUmpActivitiesGetRequest) GetToolId() int64 {
    return r.toolId
}
// PageNo Setter
// 分页的页码
func (r *TaobaoUmpActivitiesGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoUmpActivitiesGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页的最大条数
func (r *TaobaoUmpActivitiesGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoUmpActivitiesGetRequest) GetPageSize() int64 {
    return r.pageSize
}
