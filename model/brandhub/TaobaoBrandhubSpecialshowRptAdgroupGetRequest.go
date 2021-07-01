package brandhub

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌号品牌特秀单元报表数据查询 API请求
taobao.brandhub.specialshow.rpt.adgroup.get

获取品牌号品牌特秀广告adgroup分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
*/
type TaobaoBrandhubSpecialshowRptAdgroupGetAPIRequest struct {
    model.Params
    // 开始时间(最多可查询最近90天)
    _startDate   string
    // 截至时间(最晚到昨天)
    _endDate   string
    // 指定计划id
    _solutionId   string
    // 指定任务id
    _taskId   string
    // 当前页数
    _pageIndex   string
    // 可选页数
    _pageSize   string
}

// 初始化TaobaoBrandhubSpecialshowRptAdgroupGetAPIRequest对象
func NewTaobaoBrandhubSpecialshowRptAdgroupGetRequest() *TaobaoBrandhubSpecialshowRptAdgroupGetAPIRequest{
    return &TaobaoBrandhubSpecialshowRptAdgroupGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBrandhubSpecialshowRptAdgroupGetAPIRequest) GetApiMethodName() string {
    return "taobao.brandhub.specialshow.rpt.adgroup.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBrandhubSpecialshowRptAdgroupGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartDate Setter
// 开始时间(最多可查询最近90天)
func (r *TaobaoBrandhubSpecialshowRptAdgroupGetAPIRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoBrandhubSpecialshowRptAdgroupGetAPIRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 截至时间(最晚到昨天)
func (r *TaobaoBrandhubSpecialshowRptAdgroupGetAPIRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoBrandhubSpecialshowRptAdgroupGetAPIRequest) GetEndDate() string {
    return r._endDate
}
// SolutionId Setter
// 指定计划id
func (r *TaobaoBrandhubSpecialshowRptAdgroupGetAPIRequest) SetSolutionId(_solutionId string) error {
    r._solutionId = _solutionId
    r.Set("solution_id", _solutionId)
    return nil
}

// SolutionId Getter
func (r TaobaoBrandhubSpecialshowRptAdgroupGetAPIRequest) GetSolutionId() string {
    return r._solutionId
}
// TaskId Setter
// 指定任务id
func (r *TaobaoBrandhubSpecialshowRptAdgroupGetAPIRequest) SetTaskId(_taskId string) error {
    r._taskId = _taskId
    r.Set("task_id", _taskId)
    return nil
}

// TaskId Getter
func (r TaobaoBrandhubSpecialshowRptAdgroupGetAPIRequest) GetTaskId() string {
    return r._taskId
}
// PageIndex Setter
// 当前页数
func (r *TaobaoBrandhubSpecialshowRptAdgroupGetAPIRequest) SetPageIndex(_pageIndex string) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoBrandhubSpecialshowRptAdgroupGetAPIRequest) GetPageIndex() string {
    return r._pageIndex
}
// PageSize Setter
// 可选页数
func (r *TaobaoBrandhubSpecialshowRptAdgroupGetAPIRequest) SetPageSize(_pageSize string) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoBrandhubSpecialshowRptAdgroupGetAPIRequest) GetPageSize() string {
    return r._pageSize
}
