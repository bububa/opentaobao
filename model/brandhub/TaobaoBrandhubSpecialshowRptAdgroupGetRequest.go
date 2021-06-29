package brandhub

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌号品牌特秀单元报表数据查询 APIRequest
taobao.brandhub.specialshow.rpt.adgroup.get

获取品牌号品牌特秀广告adgroup分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
*/
type TaobaoBrandhubSpecialshowRptAdgroupGetRequest struct {
    model.Params

    // 开始时间(最多可查询最近90天)
    startDate   string 

    // 截至时间(最晚到昨天)
    endDate   string 

    // 指定计划id
    solutionId   string 

    // 指定任务id
    taskId   string 

    // 当前页数
    pageIndex   string 

    // 可选页数
    pageSize   string 

}

func NewTaobaoBrandhubSpecialshowRptAdgroupGetRequest() *TaobaoBrandhubSpecialshowRptAdgroupGetRequest{
    return &TaobaoBrandhubSpecialshowRptAdgroupGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBrandhubSpecialshowRptAdgroupGetRequest) GetApiMethodName() string {
    return "taobao.brandhub.specialshow.rpt.adgroup.get"
}

func (r TaobaoBrandhubSpecialshowRptAdgroupGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBrandhubSpecialshowRptAdgroupGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TaobaoBrandhubSpecialshowRptAdgroupGetRequest) GetStartDate() string {
    return r.startDate
}

func (r *TaobaoBrandhubSpecialshowRptAdgroupGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoBrandhubSpecialshowRptAdgroupGetRequest) GetEndDate() string {
    return r.endDate
}

func (r *TaobaoBrandhubSpecialshowRptAdgroupGetRequest) SetSolutionId(solutionId string) error {
    r.solutionId = solutionId
    r.Set("solution_id", solutionId)
    return nil
}

func (r TaobaoBrandhubSpecialshowRptAdgroupGetRequest) GetSolutionId() string {
    return r.solutionId
}

func (r *TaobaoBrandhubSpecialshowRptAdgroupGetRequest) SetTaskId(taskId string) error {
    r.taskId = taskId
    r.Set("task_id", taskId)
    return nil
}

func (r TaobaoBrandhubSpecialshowRptAdgroupGetRequest) GetTaskId() string {
    return r.taskId
}

func (r *TaobaoBrandhubSpecialshowRptAdgroupGetRequest) SetPageIndex(pageIndex string) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

func (r TaobaoBrandhubSpecialshowRptAdgroupGetRequest) GetPageIndex() string {
    return r.pageIndex
}

func (r *TaobaoBrandhubSpecialshowRptAdgroupGetRequest) SetPageSize(pageSize string) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoBrandhubSpecialshowRptAdgroupGetRequest) GetPageSize() string {
    return r.pageSize
}

