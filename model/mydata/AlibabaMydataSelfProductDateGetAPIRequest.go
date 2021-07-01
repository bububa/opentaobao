package mydata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMydataSelfProductDateGetAPIRequest
获取客户产品相关表现数据的可用时间范围 API请求
alibaba.mydata.self.product.date.get

获取客户产品相关表现数据的可用时间范围 */
type AlibabaMydataSelfProductDateGetAPIRequest struct {
	model.Params
	// 统计周期类型，可以为"day"，"week"，"month"
	_statisticsType string
}

// New
