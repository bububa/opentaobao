package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkOrderRefundGet 五道口订单退款按ID查询
// alibaba.wdk.order.refund.get
//
// 按照退款ID或者五道口中台订单ID查询退款信息详情
func AlibabaWdkOrderRefundGet(clt *core.SDKClient, req *wdk.AlibabaWdkOrderRefundGetAPIRequest, resp *wdk.AlibabaWdkOrderRefundGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
