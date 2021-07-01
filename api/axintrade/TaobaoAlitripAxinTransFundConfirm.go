package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

/* TaobaoAlitripAxinTransFundConfirm
确认资金单
taobao.alitrip.axin.trans.fund.confirm

通过外部订单号进行资金结算 */
func TaobaoAlitripAxinTransFundConfirm(clt *core.SDKClient, req *axintrade.TaobaoAlitripAxinTransFundConfirmAPIRequest, session string) (*axintrade.TaobaoAlitripAxinTransFundConfirmAPIResponse, error) {
	var resp axintrade.TaobaoAlitripAxinTransFundConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
