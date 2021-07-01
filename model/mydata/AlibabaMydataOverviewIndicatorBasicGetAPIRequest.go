package mydata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMydataOverviewIndicatorBasicGetAPIRequest
我的效果-获取公司询盘流量行业表现 API请求
alibaba.mydata.overview.indicator.basic.get

获取公司询盘流量行业表现 */
type AlibabaMydataOverviewIndicatorBasicGetAPIRequest struct {
	model.Params
	// 要查询的数据周期
	_dateRange *DateRange
	// 要查询的行业信息
	_industry *Industry
}

// New
