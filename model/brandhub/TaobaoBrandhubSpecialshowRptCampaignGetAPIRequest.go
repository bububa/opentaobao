package brandhub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBrandhubSpecialshowRptCampaignGetAPIRequest 品牌号品牌特秀计划报表数据查询 API请求
// taobao.brandhub.specialshow.rpt.campaign.get
//
// 获取品牌号品牌特秀广告campaign分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
type TaobaoBrandhubSpecialshowRptCampaignGetAPIRequest struct {
	model.Params
	// 开始时间(最多可查询最近90天)
	_startDate string
	// 指定计划id
	_solutionId string
	// 截至时间(最晚到昨天)
	_endDate string
	// 当前页数
	_pageIndex string
	// 每页条数
	_pageSize string
}

// NewTaobaoBrandhubSpecialshowRptCampaignGetRequest 初始化TaobaoBrandhubSpecialshowRptCampaignGetAPIRequest对象
func NewTaobaoBrandhubSpecialshowRptCampaignGetRequest() *TaobaoBrandhubSpecialshowRptCampaignGetAPIRequest {
	return &TaobaoBrandhubSpecialshowRptCampaignGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBrandhubSpecialshowRptCampaignGetAPIRequest) GetApiMethodName() string {
	return "taobao.brandhub.specialshow.rpt.campaign.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBrandhubSpecialshowRptCampaignGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStartDate is StartDate Setter
// 开始时间(最多可查询最近90天)
func (r *TaobaoBrandhubSpecialshowRptCampaignGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoBrandhubSpecialshowRptCampaignGetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetSolutionId is SolutionId Setter
// 指定计划id
func (r *TaobaoBrandhubSpecialshowRptCampaignGetAPIRequest) SetSolutionId(_solutionId string) error {
	r._solutionId = _solutionId
	r.Set("solution_id", _solutionId)
	return nil
}

// GetSolutionId SolutionId Getter
func (r TaobaoBrandhubSpecialshowRptCampaignGetAPIRequest) GetSolutionId() string {
	return r._solutionId
}

// SetEndDate is EndDate Setter
// 截至时间(最晚到昨天)
func (r *TaobaoBrandhubSpecialshowRptCampaignGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoBrandhubSpecialshowRptCampaignGetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetPageIndex is PageIndex Setter
// 当前页数
func (r *TaobaoBrandhubSpecialshowRptCampaignGetAPIRequest) SetPageIndex(_pageIndex string) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaoBrandhubSpecialshowRptCampaignGetAPIRequest) GetPageIndex() string {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *TaobaoBrandhubSpecialshowRptCampaignGetAPIRequest) SetPageSize(_pageSize string) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoBrandhubSpecialshowRptCampaignGetAPIRequest) GetPageSize() string {
	return r._pageSize
}
