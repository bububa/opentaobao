package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkOrderGet 交易订单详情查询
// alibaba.wdk.order.get
//
// 五道口三江单据查询接口
func AlibabaWdkOrderGet(clt *core.SDKClient, req *wdk.AlibabaWdkOrderGetAPIRequest, resp *wdk.AlibabaWdkOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
