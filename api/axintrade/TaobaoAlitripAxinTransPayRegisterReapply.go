package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripAxinTransPayRegisterReapply 阿信支付入驻重新申请
// taobao.alitrip.axin.trans.pay.register.reapply
//
// 阿信支付入驻重新申请
// 用于支付平台驳回，商户提交时的业务场景
func TaobaoAlitripAxinTransPayRegisterReapply(clt *core.SDKClient, req *axintrade.TaobaoAlitripAxinTransPayRegisterReapplyAPIRequest, resp *axintrade.TaobaoAlitripAxinTransPayRegisterReapplyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
