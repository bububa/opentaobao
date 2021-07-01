package mydata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMydataSelfProductGetAPIRequest
获取客户产品相关表现数据 API请求
alibaba.mydata.self.product.get

获取客户产品相关表现数据 */
type AlibabaMydataSelfProductGetAPIRequest struct {
	model.Params
	// 统计周期，可以为"day", "week", "month"
	_statisticsType string
	// 统计日期
	_statDate string
	// 待查询产品id列表
	_productIds []int64
}

// New
