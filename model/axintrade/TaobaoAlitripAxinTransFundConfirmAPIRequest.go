package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripAxinTransFundConfirmAPIRequest
确认资金单 API请求
taobao.alitrip.axin.trans.fund.confirm

通过外部订单号进行资金结算 */
type TaobaoAlitripAxinTransFundConfirmAPIRequest struct {
	model.Params
	// 外部订单编号
	_outerOrderId string
}

// New
