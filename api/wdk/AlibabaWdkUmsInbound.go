package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkUmsInbound 入库-ERP下发单
// alibaba.wdk.ums.inbound
//
// 入库-ERP下发单
func AlibabaWdkUmsInbound(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkUmsInboundAPIRequest, resp *wdk.AlibabaWdkUmsInboundAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
