package tbrefund

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRefundMessageAdd 创建退款留言/凭证
// taobao.refund.message.add
//
// 创建退款留言/凭证
func TaobaoRefundMessageAdd(ctx context.Context, clt *core.SDKClient, req *tbrefund.TaobaoRefundMessageAddAPIRequest, resp *tbrefund.TaobaoRefundMessageAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
