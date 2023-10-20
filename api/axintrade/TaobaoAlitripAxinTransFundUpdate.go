package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripAxinTransFundUpdate 修改资金单接口
// taobao.alitrip.axin.trans.fund.update
//
// 阿信供销平台-修改资金单接口
func TaobaoAlitripAxinTransFundUpdate(clt *core.SDKClient, req *axintrade.TaobaoAlitripAxinTransFundUpdateAPIRequest, resp *axintrade.TaobaoAlitripAxinTransFundUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
