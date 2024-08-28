package tbrefund

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRefundsApplyGet 查询买家申请的退款列表
// taobao.refunds.apply.get
//
// 查询买家申请的退款列表，且查询外店的退款列表时需要指定交易类型
func TaobaoRefundsApplyGet(ctx context.Context, clt *core.SDKClient, req *tbrefund.TaobaoRefundsApplyGetAPIRequest, resp *tbrefund.TaobaoRefundsApplyGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
