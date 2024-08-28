package axintrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripAxinTransPaySignCheck 阿信支付宝验签服务
// taobao.alitrip.axin.trans.pay.sign.check
//
// 阿信支付宝验签服务
func TaobaoAlitripAxinTransPaySignCheck(ctx context.Context, clt *core.SDKClient, req *axintrade.TaobaoAlitripAxinTransPaySignCheckAPIRequest, resp *axintrade.TaobaoAlitripAxinTransPaySignCheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
