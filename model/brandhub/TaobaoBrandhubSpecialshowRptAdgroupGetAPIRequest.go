package brandhub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobrandhubspecialshowrptadgroupgetAPIRequest 品牌号品牌特秀单元报表数据查询 API请求
// taobao.brandhub.specialshow.rpt.adgroup.get
//
// 获取品牌号品牌特秀广告adgroup分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
type TaobaobrandhubspecialshowrptadgroupgetAPIRequest struct {
	model.Params
	// 开始时间(最多可查询最近90天)
	_startDate string
	// 指定计划id
	_solutionId string
	// 指定任务id
	_taskId string
	// 截至时间(最晚到昨天)
	_endDate string
	// 当前页数
	_pageIndex string
	// 可选页数
	_pageSize string
}

// NewTaobaobrandhubspecialshowrptadgroupgetRequest 初始化TaobaobrandhubspecialshowrptadgroupgetAPIRequest对象
func NewTaobaobrandhubspecialshowrptadgroupgetRequest() *TaobaobrandhubspecialshowrptadgroupgetAPIRequest {
	return &TaobaobrandhubspecialshowrptadgroupgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobrandhubspecialshowrptadgroupgetAPIRequest) GetApiMethodName() string {
	return "taobao.brandhub.specialshow.rpt.adgroup.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobrandhubspecialshowrptadgroupgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobrandhubspecialshowrptadgroupgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartDate is StartDate Setter
// 开始时间(最多可查询最近90天)
func (r *TaobaobrandhubspecialshowrptadgroupgetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaobrandhubspecialshowrptadgroupgetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetSolutionId is SolutionId Setter
// 指定计划id
func (r *TaobaobrandhubspecialshowrptadgroupgetAPIRequest) SetSolutionId(_solutionId string) error {
	r._solutionId = _solutionId
	r.Set("solution_id", _solutionId)
	return nil
}

// GetSolutionId SolutionId Getter
func (r TaobaobrandhubspecialshowrptadgroupgetAPIRequest) GetSolutionId() string {
	return r._solutionId
}

// SetTaskId is TaskId Setter
// 指定任务id
func (r *TaobaobrandhubspecialshowrptadgroupgetAPIRequest) SetTaskId(_taskId string) error {
	r._taskId = _taskId
	r.Set("task_id", _taskId)
	return nil
}

// GetTaskId TaskId Getter
func (r TaobaobrandhubspecialshowrptadgroupgetAPIRequest) GetTaskId() string {
	return r._taskId
}

// SetEndDate is EndDate Setter
// 截至时间(最晚到昨天)
func (r *TaobaobrandhubspecialshowrptadgroupgetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaobrandhubspecialshowrptadgroupgetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetPageIndex is PageIndex Setter
// 当前页数
func (r *TaobaobrandhubspecialshowrptadgroupgetAPIRequest) SetPageIndex(_pageIndex string) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaobrandhubspecialshowrptadgroupgetAPIRequest) GetPageIndex() string {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 可选页数
func (r *TaobaobrandhubspecialshowrptadgroupgetAPIRequest) SetPageSize(_pageSize string) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaobrandhubspecialshowrptadgroupgetAPIRequest) GetPageSize() string {
	return r._pageSize
}
