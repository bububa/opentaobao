package xhotelofficial

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelofficial"
)

// TaobaoXhotelOrderOfficialCancel 官网信用住订单取消
// taobao.xhotel.order.official.cancel
//
// 官网信用住订单取消
func TaobaoXhotelOrderOfficialCancel(ctx context.Context, clt *core.SDKClient, req *xhotelofficial.TaobaoXhotelOrderOfficialCancelAPIRequest, resp *xhotelofficial.TaobaoXhotelOrderOfficialCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
