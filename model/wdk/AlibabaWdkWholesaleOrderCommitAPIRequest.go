package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkWholesaleOrderCommitAPIRequest
盒马帮采购确认订单接口 API请求
alibaba.wdk.wholesale.order.commit

盒马帮采购确认订单接口 */
type AlibabaWdkWholesaleOrderCommitAPIRequest struct {
	model.Params
	// 采购单信息
	_orderCommitReq *OrderCommitReq
}

// New
