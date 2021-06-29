package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取活动列表接口 APIRequest
alibaba.westcrm.activity.list.get

获取活动列表提供给友盟&互动屏
*/
type AlibabaWestcrmActivityListGetRequest struct {
    model.Params

    // 活动状态
    status   int64 

    // 园区id
    campusId   int64 

    // 排序方向
    sord   string 

    // 页,默认第一页
    page   int64 

    // 排序字段
    sidx   string 

    // 分页偏移量 eq . limit offset ,rows
    offset   int64 

    // 页大小,默认每页查询10条数据
    rows   int64 

}

func NewAlibabaWestcrmActivityListGetRequest() *AlibabaWestcrmActivityListGetRequest{
    return &AlibabaWestcrmActivityListGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWestcrmActivityListGetRequest) GetApiMethodName() string {
    return "alibaba.westcrm.activity.list.get"
}

func (r AlibabaWestcrmActivityListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWestcrmActivityListGetRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaWestcrmActivityListGetRequest) GetStatus() int64 {
    return r.status
}

func (r *AlibabaWestcrmActivityListGetRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaWestcrmActivityListGetRequest) GetCampusId() int64 {
    return r.campusId
}

func (r *AlibabaWestcrmActivityListGetRequest) SetSord(sord string) error {
    r.sord = sord
    r.Set("sord", sord)
    return nil
}

func (r AlibabaWestcrmActivityListGetRequest) GetSord() string {
    return r.sord
}

func (r *AlibabaWestcrmActivityListGetRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaWestcrmActivityListGetRequest) GetPage() int64 {
    return r.page
}

func (r *AlibabaWestcrmActivityListGetRequest) SetSidx(sidx string) error {
    r.sidx = sidx
    r.Set("sidx", sidx)
    return nil
}

func (r AlibabaWestcrmActivityListGetRequest) GetSidx() string {
    return r.sidx
}

func (r *AlibabaWestcrmActivityListGetRequest) SetOffset(offset int64) error {
    r.offset = offset
    r.Set("offset", offset)
    return nil
}

func (r AlibabaWestcrmActivityListGetRequest) GetOffset() int64 {
    return r.offset
}

func (r *AlibabaWestcrmActivityListGetRequest) SetRows(rows int64) error {
    r.rows = rows
    r.Set("rows", rows)
    return nil
}

func (r AlibabaWestcrmActivityListGetRequest) GetRows() int64 {
    return r.rows
}

