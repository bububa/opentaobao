package tbrefund

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoSpecialRefundsReceiveGet 特殊退款类型的纠纷单列表查询
// taobao.special.refunds.receive.get
//
// 特殊退款类型的纠纷单列表查询
func TaobaoSpecialRefundsReceiveGet(ctx context.Context, clt *core.SDKClient, req *tbrefund.TaobaoSpecialRefundsReceiveGetAPIRequest, resp *tbrefund.TaobaoSpecialRefundsReceiveGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
