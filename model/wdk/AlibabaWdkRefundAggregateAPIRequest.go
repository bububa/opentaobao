package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkRefundAggregateAPIRequest
淘鲜达退款单按门店聚合查询 API请求
alibaba.wdk.refund.aggregate

淘鲜达退款单按门店聚合查询 */
type AlibabaWdkRefundAggregateAPIRequest struct {
	model.Params
	// 系统自动生成
	_refundAggregateQueryRequest *RefundAggregateQueryRequest
}

// New
