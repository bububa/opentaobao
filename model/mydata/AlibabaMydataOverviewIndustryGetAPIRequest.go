package mydata

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMydataOverviewIndustryGetAPIRequest) Reset() {
	r._dateRange = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMydataOverviewIndustryGetAPIRequest) GetApiMethodName() string {
	return "alibaba.mydata.overview.industry.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMydataOverviewIndustryGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMydataOverviewIndustryGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDateRange is DateRange Setter
// 系统自动生成
func (r *AlibabaMydataOverviewIndustryGetAPIRequest) SetDateRange(_dateRange *DateRange) error {
	r._dateRange = _dateRange
	r.Set("date_range", _dateRange)
	return nil
}

// GetDateRange DateRange Getter
func (r AlibabaMydataOverviewIndustryGetAPIRequest) GetDateRange() *DateRange {
	return r._dateRange
}

var poolAlibabaMydataOverviewIndustryGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMydataOverviewIndustryGetRequest()
	},
}

// GetAlibabaMydataOverviewIndustryGetRequest 从 sync.Pool 获取 AlibabaMydataOverviewIndustryGetAPIRequest
func GetAlibabaMydataOverviewIndustryGetAPIRequest() *AlibabaMydataOverviewIndustryGetAPIRequest {
	return poolAlibabaMydataOverviewIndustryGetAPIRequest.Get().(*AlibabaMydataOverviewIndustryGetAPIRequest)
}

// ReleaseAlibabaMydataOverviewIndustryGetAPIRequest 将 AlibabaMydataOverviewIndustryGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaMydataOverviewIndustryGetAPIRequest(v *AlibabaMydataOverviewIndustryGetAPIRequest) {
	v.Reset()
	poolAlibabaMydataOverviewIndustryGetAPIRequest.Put(v)
}
