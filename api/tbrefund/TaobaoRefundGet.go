package tbrefund

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRefundGet 获取单笔退款详情
// taobao.refund.get
//
// 获取单笔退款详情
func TaobaoRefundGet(ctx context.Context, clt *core.SDKClient, req *tbrefund.TaobaoRefundGetAPIRequest, resp *tbrefund.TaobaoRefundGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
