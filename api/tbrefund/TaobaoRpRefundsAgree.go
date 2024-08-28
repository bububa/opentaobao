package tbrefund

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRpRefundsAgree 同意退款
// taobao.rp.refunds.agree
//
// 卖家同意退款，支持批量退款，只允许子账号操作。淘宝退款一次最多能退20笔，总金额不超过6000元；天猫退款一次最多能退30笔，总金额不超过10000元。
func TaobaoRpRefundsAgree(ctx context.Context, clt *core.SDKClient, req *tbrefund.TaobaoRpRefundsAgreeAPIRequest, resp *tbrefund.TaobaoRpRefundsAgreeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
