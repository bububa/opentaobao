package mydata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamydataoverviewindicatorbasicgetAPIRequest 我的效果-获取公司询盘流量行业表现 API请求
// alibaba.mydata.overview.indicator.basic.get
//
// 获取公司询盘流量行业表现
type AlibabamydataoverviewindicatorbasicgetAPIRequest struct {
	model.Params
	// 要查询的数据周期
	_dateRange *DateRange
	// 要查询的行业信息
	_industry *Industry
}

// NewAlibabamydataoverviewindicatorbasicgetRequest 初始化AlibabamydataoverviewindicatorbasicgetAPIRequest对象
func NewAlibabamydataoverviewindicatorbasicgetRequest() *AlibabamydataoverviewindicatorbasicgetAPIRequest {
	return &AlibabamydataoverviewindicatorbasicgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamydataoverviewindicatorbasicgetAPIRequest) GetApiMethodName() string {
	return "alibaba.mydata.overview.indicator.basic.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamydataoverviewindicatorbasicgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamydataoverviewindicatorbasicgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDateRange is DateRange Setter
// 要查询的数据周期
func (r *AlibabamydataoverviewindicatorbasicgetAPIRequest) SetDateRange(_dateRange *DateRange) error {
	r._dateRange = _dateRange
	r.Set("date_range", _dateRange)
	return nil
}

// GetDateRange DateRange Getter
func (r AlibabamydataoverviewindicatorbasicgetAPIRequest) GetDateRange() *DateRange {
	return r._dateRange
}

// SetIndustry is Industry Setter
// 要查询的行业信息
func (r *AlibabamydataoverviewindicatorbasicgetAPIRequest) SetIndustry(_industry *Industry) error {
	r._industry = _industry
	r.Set("industry", _industry)
	return nil
}

// GetIndustry Industry Getter
func (r AlibabamydataoverviewindicatorbasicgetAPIRequest) GetIndustry() *Industry {
	return r._industry
}
