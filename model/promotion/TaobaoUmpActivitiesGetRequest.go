package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询活动列表 APIRequest
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

func NewTaobaoUmpActivitiesGetRequest() *TaobaoUmpActivitiesGetRequest{
    return &TaobaoUmpActivitiesGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpActivitiesGetRequest) GetApiMethodName() string {
    return "taobao.ump.activities.get"
}

func (r TaobaoUmpActivitiesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUmpActivitiesGetRequest) SetToolId(toolId int64) error {
    r.toolId = toolId
    r.Set("tool_id", toolId)
    return nil
}

func (r TaobaoUmpActivitiesGetRequest) GetToolId() int64 {
    return r.toolId
}

func (r *TaobaoUmpActivitiesGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoUmpActivitiesGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoUmpActivitiesGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoUmpActivitiesGetRequest) GetPageSize() int64 {
    return r.pageSize
}

