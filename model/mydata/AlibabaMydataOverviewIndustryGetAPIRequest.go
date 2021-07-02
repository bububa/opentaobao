package mydata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMydataOverviewIndustryGetAPIRequest 我的效果-获取Top行业列表 API请求
// alibaba.mydata.overview.industry.get
//
// 获取数据管家我的效果API可以使用的行业
type AlibabaMydataOverviewIndustryGetAPIRequest struct {
	model.Params
	// 系统自动生成
	_dateRange *DateRange
}

// NewAlibabaMydataOverviewIndustryGetRequest 初始化AlibabaMydataOverviewIndustryGetAPIRequest对象
func NewAlibabaMydataOverviewIndustryGetRequest() *AlibabaMydataOverviewIndustryGetAPIRequest {
	return &AlibabaMydataOverviewIndustryGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMydataOverviewIndustryGetAPIRequest) GetApiMethodName() string {
	return "alibaba.mydata.overview.industry.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMydataOverviewIndustryGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DateRange Setter
// 系统自动生成
func (r *AlibabaMydataOverviewIndustryGetAPIRequest) SetDateRange(_dateRange *DateRange) error {
	r._dateRange = _dateRange
	r.Set("date_range", _dateRange)
	return nil
}

// Get DateRange Getter
func (r AlibabaMydataOverviewIndustryGetAPIRequest) GetDateRange() *DateRange {
	return r._dateRange
}
