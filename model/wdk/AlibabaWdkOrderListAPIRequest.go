package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkOrderListAPIRequest
五道口订单拉取 API请求
alibaba.wdk.order.list

五道口交易订单拉取接口 */
type AlibabaWdkOrderListAPIRequest struct {
	model.Params
	// 查询参数
	_batchQueryRequest *BatchQueryRequest
}

// New
