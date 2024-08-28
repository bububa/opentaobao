package icbuassurance

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuassurance"
)

// AlibabaIcbuTradeAssuranceAccountGet icbu信保账户信息
// alibaba.icbu.trade.assurance.account.get
//
// icbu交易信用保障开通状态&amp;额度信息查询
func AlibabaIcbuTradeAssuranceAccountGet(ctx context.Context, clt *core.SDKClient, req *icbuassurance.AlibabaIcbuTradeAssuranceAccountGetAPIRequest, resp *icbuassurance.AlibabaIcbuTradeAssuranceAccountGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
