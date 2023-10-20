package mydata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamydataoverviewindustrygetAPIRequest 我的效果-获取Top行业列表 API请求
// alibaba.mydata.overview.industry.get
//
// 获取数据管家我的效果API可以使用的行业
type AlibabamydataoverviewindustrygetAPIRequest struct {
	model.Params
	// 系统自动生成
	_dateRange *DateRange
}

// NewAlibabamydataoverviewindustrygetRequest 初始化AlibabamydataoverviewindustrygetAPIRequest对象
func NewAlibabamydataoverviewindustrygetRequest() *AlibabamydataoverviewindustrygetAPIRequest {
	return &AlibabamydataoverviewindustrygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamydataoverviewindustrygetAPIRequest) GetApiMethodName() string {
	return "alibaba.mydata.overview.industry.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamydataoverviewindustrygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamydataoverviewindustrygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDateRange is DateRange Setter
// 系统自动生成
func (r *AlibabamydataoverviewindustrygetAPIRequest) SetDateRange(_dateRange *DateRange) error {
	r._dateRange = _dateRange
	r.Set("date_range", _dateRange)
	return nil
}

// GetDateRange DateRange Getter
func (r AlibabamydataoverviewindustrygetAPIRequest) GetDateRange() *DateRange {
	return r._dateRange
}
