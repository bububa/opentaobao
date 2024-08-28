package axintrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripAxinTransPayRegisterAudit 阿信支付入驻审核通知
// taobao.alitrip.axin.trans.pay.register.audit
//
// 阿信支付入驻审核通知
func TaobaoAlitripAxinTransPayRegisterAudit(ctx context.Context, clt *core.SDKClient, req *axintrade.TaobaoAlitripAxinTransPayRegisterAuditAPIRequest, resp *axintrade.TaobaoAlitripAxinTransPayRegisterAuditAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
