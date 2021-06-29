package brandhub

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌号品牌特秀计划报表数据查询 API请求
taobao.brandhub.specialshow.rpt.campaign.get

获取品牌号品牌特秀广告campaign分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
*/
type TaobaoBrandhubSpecialshowRptCampaignGetRequest struct {
    model.Params
    // 开始时间(最多可查询最近90天)
    startDate   string
    // 指定计划id
    solutionId   string
    // 截至时间(最晚到昨天)
    endDate   string
    // 当前页数
    pageIndex   string
    // 每页条数
    pageSize   string
}

// 初始化TaobaoBrandhubSpecialshowRptCampaignGetRequest对象
func NewTaobaoBrandhubSpecialshowRptCampaignGetRequest() *TaobaoBrandhubSpecialshowRptCampaignGetRequest{
    return &TaobaoBrandhubSpecialshowRptCampaignGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBrandhubSpecialshowRptCampaignGetRequest) GetApiMethodName() string {
    return "taobao.brandhub.specialshow.rpt.campaign.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBrandhubSpecialshowRptCampaignGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartDate Setter
// 开始时间(最多可查询最近90天)
func (r *TaobaoBrandhubSpecialshowRptCampaignGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TaobaoBrandhubSpecialshowRptCampaignGetRequest) GetStartDate() string {
    return r.startDate
}
// SolutionId Setter
// 指定计划id
func (r *TaobaoBrandhubSpecialshowRptCampaignGetRequest) SetSolutionId(solutionId string) error {
    r.solutionId = solutionId
    r.Set("solution_id", solutionId)
    return nil
}

// SolutionId Getter
func (r TaobaoBrandhubSpecialshowRptCampaignGetRequest) GetSolutionId() string {
    return r.solutionId
}
// EndDate Setter
// 截至时间(最晚到昨天)
func (r *TaobaoBrandhubSpecialshowRptCampaignGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoBrandhubSpecialshowRptCampaignGetRequest) GetEndDate() string {
    return r.endDate
}
// PageIndex Setter
// 当前页数
func (r *TaobaoBrandhubSpecialshowRptCampaignGetRequest) SetPageIndex(pageIndex string) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoBrandhubSpecialshowRptCampaignGetRequest) GetPageIndex() string {
    return r.pageIndex
}
// PageSize Setter
// 每页条数
func (r *TaobaoBrandhubSpecialshowRptCampaignGetRequest) SetPageSize(pageSize string) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoBrandhubSpecialshowRptCampaignGetRequest) GetPageSize() string {
    return r.pageSize
}
