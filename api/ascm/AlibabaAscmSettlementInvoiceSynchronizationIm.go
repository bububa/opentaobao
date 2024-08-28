package ascm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascm"
)

// AlibabaAscmSettlementInvoiceSynchronizationIm 英迈发票同步到结算
// alibaba.ascm.settlement.invoice.synchronization.im
//
// 外部供应商通过此API将发货的发票信息推送给供应链中台结算系统
func AlibabaAscmSettlementInvoiceSynchronizationIm(ctx context.Context, clt *core.SDKClient, req *ascm.AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest, resp *ascm.AlibabaAscmSettlementInvoiceSynchronizationImAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
