package axintrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripAxinTransPayRegisterCreate 提交支付服务开通
// taobao.alitrip.axin.trans.pay.register.create
//
// 阿信供销平台-提交支付服务开通接口
func TaobaoAlitripAxinTransPayRegisterCreate(ctx context.Context, clt *core.SDKClient, req *axintrade.TaobaoAlitripAxinTransPayRegisterCreateAPIRequest, resp *axintrade.TaobaoAlitripAxinTransPayRegisterCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
