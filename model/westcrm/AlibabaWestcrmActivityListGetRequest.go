package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取活动列表接口 API请求
alibaba.westcrm.activity.list.get

获取活动列表提供给友盟&互动屏
*/
type AlibabaWestcrmActivityListGetRequest struct {
    model.Params
    // 活动状态
    _status   int64
    // 园区id
    _campusId   int64
    // 排序方向
    _sord   string
    // 页,默认第一页
    _page   int64
    // 排序字段
    _sidx   string
    // 分页偏移量 eq . limit offset ,rows
    _offset   int64
    // 页大小,默认每页查询10条数据
    _rows   int64
}

// 初始化AlibabaWestcrmActivityListGetRequest对象
func NewAlibabaWestcrmActivityListGetRequest() *AlibabaWestcrmActivityListGetRequest{
    return &AlibabaWestcrmActivityListGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmActivityListGetRequest) GetApiMethodName() string {
    return "alibaba.westcrm.activity.list.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmActivityListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Status Setter
// 活动状态
func (r *AlibabaWestcrmActivityListGetRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaWestcrmActivityListGetRequest) GetStatus() int64 {
    return r._status
}
// CampusId Setter
// 园区id
func (r *AlibabaWestcrmActivityListGetRequest) SetCampusId(_campusId int64) error {
    r._campusId = _campusId
    r.Set("campus_id", _campusId)
    return nil
}

// CampusId Getter
func (r AlibabaWestcrmActivityListGetRequest) GetCampusId() int64 {
    return r._campusId
}
// Sord Setter
// 排序方向
func (r *AlibabaWestcrmActivityListGetRequest) SetSord(_sord string) error {
    r._sord = _sord
    r.Set("sord", _sord)
    return nil
}

// Sord Getter
func (r AlibabaWestcrmActivityListGetRequest) GetSord() string {
    return r._sord
}
// Page Setter
// 页,默认第一页
func (r *AlibabaWestcrmActivityListGetRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaWestcrmActivityListGetRequest) GetPage() int64 {
    return r._page
}
// Sidx Setter
// 排序字段
func (r *AlibabaWestcrmActivityListGetRequest) SetSidx(_sidx string) error {
    r._sidx = _sidx
    r.Set("sidx", _sidx)
    return nil
}

// Sidx Getter
func (r AlibabaWestcrmActivityListGetRequest) GetSidx() string {
    return r._sidx
}
// Offset Setter
// 分页偏移量 eq . limit offset ,rows
func (r *AlibabaWestcrmActivityListGetRequest) SetOffset(_offset int64) error {
    r._offset = _offset
    r.Set("offset", _offset)
    return nil
}

// Offset Getter
func (r AlibabaWestcrmActivityListGetRequest) GetOffset() int64 {
    return r._offset
}
// Rows Setter
// 页大小,默认每页查询10条数据
func (r *AlibabaWestcrmActivityListGetRequest) SetRows(_rows int64) error {
    r._rows = _rows
    r.Set("rows", _rows)
    return nil
}

// Rows Getter
func (r AlibabaWestcrmActivityListGetRequest) GetRows() int64 {
    return r._rows
}
