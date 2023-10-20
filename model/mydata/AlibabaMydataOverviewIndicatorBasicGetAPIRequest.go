package mydata

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMydataOverviewIndicatorBasicGetAPIRequest 我的效果-获取公司询盘流量行业表现 API请求
// alibaba.mydata.overview.indicator.basic.get
//
// 获取公司询盘流量行业表现
type AlibabaMydataOverviewIndicatorBasicGetAPIRequest struct {
	model.Params
	// 要查询的数据周期
	_dateRange *DateRange
	// 要查询的行业信息
	_industry *Industry
}

// NewAlibabaMydataOverviewIndicatorBasicGetRequest 初始化AlibabaMydataOverviewIndicatorBasicGetAPIRequest对象
func NewAlibabaMydataOverviewIndicatorBasicGetRequest() *AlibabaMydataOverviewIndicatorBasicGetAPIRequest {
	return &AlibabaMydataOverviewIndicatorBasicGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMydataOverviewIndicatorBasicGetAPIRequest) Reset() {
	r._dateRange = nil
	r._industry = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMydataOverviewIndicatorBasicGetAPIRequest) GetApiMethodName() string {
	return "alibaba.mydata.overview.indicator.basic.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMydataOverviewIndicatorBasicGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMydataOverviewIndicatorBasicGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDateRange is DateRange Setter
// 要查询的数据周期
func (r *AlibabaMydataOverviewIndicatorBasicGetAPIRequest) SetDateRange(_dateRange *DateRange) error {
	r._dateRange = _dateRange
	r.Set("date_range", _dateRange)
	return nil
}

// GetDateRange DateRange Getter
func (r AlibabaMydataOverviewIndicatorBasicGetAPIRequest) GetDateRange() *DateRange {
	return r._dateRange
}

// SetIndustry is Industry Setter
// 要查询的行业信息
func (r *AlibabaMydataOverviewIndicatorBasicGetAPIRequest) SetIndustry(_industry *Industry) error {
	r._industry = _industry
	r.Set("industry", _industry)
	return nil
}

// GetIndustry Industry Getter
func (r AlibabaMydataOverviewIndicatorBasicGetAPIRequest) GetIndustry() *Industry {
	return r._industry
}

var poolAlibabaMydataOverviewIndicatorBasicGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMydataOverviewIndicatorBasicGetRequest()
	},
}

// GetAlibabaMydataOverviewIndicatorBasicGetRequest 从 sync.Pool 获取 AlibabaMydataOverviewIndicatorBasicGetAPIRequest
func GetAlibabaMydataOverviewIndicatorBasicGetAPIRequest() *AlibabaMydataOverviewIndicatorBasicGetAPIRequest {
	return poolAlibabaMydataOverviewIndicatorBasicGetAPIRequest.Get().(*AlibabaMydataOverviewIndicatorBasicGetAPIRequest)
}

// ReleaseAlibabaMydataOverviewIndicatorBasicGetAPIRequest 将 AlibabaMydataOverviewIndicatorBasicGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaMydataOverviewIndicatorBasicGetAPIRequest(v *AlibabaMydataOverviewIndicatorBasicGetAPIRequest) {
	v.Reset()
	poolAlibabaMydataOverviewIndicatorBasicGetAPIRequest.Put(v)
}
