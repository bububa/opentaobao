package tbrefund

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoSpecialRefundGet 特殊部分退纠纷单查询
// taobao.special.refund.get
//
// 获取单笔特殊部分退的纠纷单查询
func TaobaoSpecialRefundGet(ctx context.Context, clt *core.SDKClient, req *tbrefund.TaobaoSpecialRefundGetAPIRequest, resp *tbrefund.TaobaoSpecialRefundGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
