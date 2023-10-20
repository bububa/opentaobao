package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkUmsOutboundProcessGet 出库业务UMS异步处理结果返回
// alibaba.wdk.ums.outbound.process.get
//
// 出库业务UMS异步处理结果返回
func AlibabaWdkUmsOutboundProcessGet(clt *core.SDKClient, req *wdk.AlibabaWdkUmsOutboundProcessGetAPIRequest, resp *wdk.AlibabaWdkUmsOutboundProcessGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
