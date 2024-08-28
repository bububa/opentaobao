package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkUmsOutbound 出库-ERP下发单(新接口，包含调拨出库单和退货出库单等)
// alibaba.wdk.ums.outbound
//
// 出库-ERP下发单(新接口，包含调拨出库单和退货出库单等)
func AlibabaWdkUmsOutbound(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkUmsOutboundAPIRequest, resp *wdk.AlibabaWdkUmsOutboundAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
