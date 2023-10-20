package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkOrderRefundList 五道口交易退款批量查询
// alibaba.wdk.order.refund.list
//
// 按照条件查询退款数据
func AlibabaWdkOrderRefundList(clt *core.SDKClient, req *wdk.AlibabaWdkOrderRefundListAPIRequest, resp *wdk.AlibabaWdkOrderRefundListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
