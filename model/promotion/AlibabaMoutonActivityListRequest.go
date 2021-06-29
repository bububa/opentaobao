package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家自运营活动列表 API请求
alibaba.mouton.activity.list

商家查询自己配置的活动列表
*/
type AlibabaMoutonActivityListRequest struct {
    model.Params
    // 开始时间
    startTimeEnd   string
    // 每页记录数
    pageSize   int64
    // 来源
    source   string
    // 开始时间
    startTimeBegin   string
    // 结束时间
    endTimeBegin   string
    // 结束时间
    endTimeEnd   string
    // 来源记录id
    sourceRecordId   int64
    // 状态
    statusList   []string
    // 当前页
    currentPage   int64
}

// 初始化AlibabaMoutonActivityListRequest对象
func NewAlibabaMoutonActivityListRequest() *AlibabaMoutonActivityListRequest{
    return &AlibabaMoutonActivityListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoutonActivityListRequest) GetApiMethodName() string {
    return "alibaba.mouton.activity.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoutonActivityListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartTimeEnd Setter
// 开始时间
func (r *AlibabaMoutonActivityListRequest) SetStartTimeEnd(startTimeEnd string) error {
    r.startTimeEnd = startTimeEnd
    r.Set("start_time_end", startTimeEnd)
    return nil
}

// StartTimeEnd Getter
func (r AlibabaMoutonActivityListRequest) GetStartTimeEnd() string {
    return r.startTimeEnd
}
// PageSize Setter
// 每页记录数
func (r *AlibabaMoutonActivityListRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaMoutonActivityListRequest) GetPageSize() int64 {
    return r.pageSize
}
// Source Setter
// 来源
func (r *AlibabaMoutonActivityListRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r AlibabaMoutonActivityListRequest) GetSource() string {
    return r.source
}
// StartTimeBegin Setter
// 开始时间
func (r *AlibabaMoutonActivityListRequest) SetStartTimeBegin(startTimeBegin string) error {
    r.startTimeBegin = startTimeBegin
    r.Set("start_time_begin", startTimeBegin)
    return nil
}

// StartTimeBegin Getter
func (r AlibabaMoutonActivityListRequest) GetStartTimeBegin() string {
    return r.startTimeBegin
}
// EndTimeBegin Setter
// 结束时间
func (r *AlibabaMoutonActivityListRequest) SetEndTimeBegin(endTimeBegin string) error {
    r.endTimeBegin = endTimeBegin
    r.Set("end_time_begin", endTimeBegin)
    return nil
}

// EndTimeBegin Getter
func (r AlibabaMoutonActivityListRequest) GetEndTimeBegin() string {
    return r.endTimeBegin
}
// EndTimeEnd Setter
// 结束时间
func (r *AlibabaMoutonActivityListRequest) SetEndTimeEnd(endTimeEnd string) error {
    r.endTimeEnd = endTimeEnd
    r.Set("end_time_end", endTimeEnd)
    return nil
}

// EndTimeEnd Getter
func (r AlibabaMoutonActivityListRequest) GetEndTimeEnd() string {
    return r.endTimeEnd
}
// SourceRecordId Setter
// 来源记录id
func (r *AlibabaMoutonActivityListRequest) SetSourceRecordId(sourceRecordId int64) error {
    r.sourceRecordId = sourceRecordId
    r.Set("source_record_id", sourceRecordId)
    return nil
}

// SourceRecordId Getter
func (r AlibabaMoutonActivityListRequest) GetSourceRecordId() int64 {
    return r.sourceRecordId
}
// StatusList Setter
// 状态
func (r *AlibabaMoutonActivityListRequest) SetStatusList(statusList []string) error {
    r.statusList = statusList
    r.Set("status_list", statusList)
    return nil
}

// StatusList Getter
func (r AlibabaMoutonActivityListRequest) GetStatusList() []string {
    return r.statusList
}
// CurrentPage Setter
// 当前页
func (r *AlibabaMoutonActivityListRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

// CurrentPage Getter
func (r AlibabaMoutonActivityListRequest) GetCurrentPage() int64 {
    return r.currentPage
}
