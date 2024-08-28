package xhotelofficial

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelofficial"
)

// TaobaoXhotelOrderOfficialSettleCancel 官网信用住取消结账
// taobao.xhotel.order.official.settle.cancel
//
// 用于官网信用住取消结账
func TaobaoXhotelOrderOfficialSettleCancel(ctx context.Context, clt *core.SDKClient, req *xhotelofficial.TaobaoXhotelOrderOfficialSettleCancelAPIRequest, resp *xhotelofficial.TaobaoXhotelOrderOfficialSettleCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
