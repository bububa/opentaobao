package mydata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMydataOverviewIndustryGetAPIRequest
我的效果-获取Top行业列表 API请求
alibaba.mydata.overview.industry.get

获取数据管家我的效果API可以使用的行业 */
type AlibabaMydataOverviewIndustryGetAPIRequest struct {
	model.Params
	// 系统自动生成
	_dateRange *DateRange
}

// New
