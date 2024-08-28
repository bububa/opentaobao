package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkUmsOutboundProcessGet 出库业务UMS异步处理结果返回
// alibaba.wdk.ums.outbound.process.get
//
// 出库业务UMS异步处理结果返回
func AlibabaWdkUmsOutboundProcessGet(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkUmsOutboundProcessGetAPIRequest, resp *wdk.AlibabaWdkUmsOutboundProcessGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
