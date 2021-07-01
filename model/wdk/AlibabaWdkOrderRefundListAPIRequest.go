package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkOrderRefundListAPIRequest
五道口交易退款批量查询 API请求
alibaba.wdk.order.refund.list

按照条件查询退款数据 */
type AlibabaWdkOrderRefundListAPIRequest struct {
	model.Params
	// 查询条件
	_batchQueryRefundRequest *BatchQueryRefundRequest
}

// New
