package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeCloseAPIRequest
卖家关闭一笔交易 API请求
taobao.trade.close

关闭一笔订单，可以是主订单或子订单。当订单从创建到关闭时间小于10s的时候，会报“CLOSE_TRADE_TOO_FAST”错误。 */
type TaobaoTradeCloseAPIRequest struct {
	model.Params
	// 主订单或子订单编号。
	_tid int64
	// 交易关闭原因。可以选择的理由有：1.未及时付款2、买家不想买了3、买家信息填写错误，重新拍4、恶意买家/同行捣乱5、缺货6、买家拍错了7、同城见面交易
	_closeReason string
}

// New
