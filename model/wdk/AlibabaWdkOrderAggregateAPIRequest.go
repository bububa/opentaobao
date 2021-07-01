package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkOrderAggregateAPIRequest
淘鲜达订单按门店机台号聚合查询 API请求
alibaba.wdk.order.aggregate

淘鲜达订单按门店机台号聚合查询 */
type AlibabaWdkOrderAggregateAPIRequest struct {
	model.Params
	// 系统自动生成
	_orderAggregateQueryRequest *OrderAggregateQueryRequest
}

// New
