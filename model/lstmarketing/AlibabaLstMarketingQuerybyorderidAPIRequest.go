package lstmarketing

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstMarketingQuerybyorderidAPIRequest
根据订单查询营销信息 API请求
alibaba.lst.marketing.querybyorderid

根据订单查询营销信息 */
type AlibabaLstMarketingQuerybyorderidAPIRequest struct {
	model.Params
	// 主订单
	_mainOrderId int64
	// 子订单
	_subOrderId int64
}

// New
