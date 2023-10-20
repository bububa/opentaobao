package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripAxinTransPaySignCheck 阿信支付宝验签服务
// taobao.alitrip.axin.trans.pay.sign.check
//
// 阿信支付宝验签服务
func TaobaoAlitripAxinTransPaySignCheck(clt *core.SDKClient, req *axintrade.TaobaoAlitripAxinTransPaySignCheckAPIRequest, resp *axintrade.TaobaoAlitripAxinTransPaySignCheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
