package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripAxinTransFundAdd 创建资金单接口
// taobao.alitrip.axin.trans.fund.add
//
// 创建资金单
func TaobaoAlitripAxinTransFundAdd(clt *core.SDKClient, req *axintrade.TaobaoAlitripAxinTransFundAddAPIRequest, resp *axintrade.TaobaoAlitripAxinTransFundAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
