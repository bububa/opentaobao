package mydata

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
我的效果-获取公司询盘流量行业表现 API请求
alibaba.mydata.overview.indicator.basic.get

获取公司询盘流量行业表现
*/
type AlibabaMydataOverviewIndicatorBasicGetAPIRequest struct {
    model.Params
    // 要查询的数据周期
    _dateRange   *DateRange
    // 要查询的行业信息
    _industry   *Industry
}

// 初始化AlibabaMydataOverviewIndicatorBasicGetAPIRequest对象
func NewAlibabaMydataOverviewIndicatorBasicGetRequest() *AlibabaMydataOverviewIndicatorBasicGetAPIRequest{
    return &AlibabaMydataOverviewIndicatorBasicGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMydataOverviewIndicatorBasicGetAPIRequest) GetApiMethodName() string {
    return "alibaba.mydata.overview.indicator.basic.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMydataOverviewIndicatorBasicGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DateRange Setter
// 要查询的数据周期
func (r *AlibabaMydataOverviewIndicatorBasicGetAPIRequest) SetDateRange(_dateRange *DateRange) error {
    r._dateRange = _dateRange
    r.Set("date_range", _dateRange)
    return nil
}

// DateRange Getter
func (r AlibabaMydataOverviewIndicatorBasicGetAPIRequest) GetDateRange() *DateRange {
    return r._dateRange
}
// Industry Setter
// 要查询的行业信息
func (r *AlibabaMydataOverviewIndicatorBasicGetAPIRequest) SetIndustry(_industry *Industry) error {
    r._industry = _industry
    r.Set("industry", _industry)
    return nil
}

// Industry Getter
func (r AlibabaMydataOverviewIndicatorBasicGetAPIRequest) GetIndustry() *Industry {
    return r._industry
}
