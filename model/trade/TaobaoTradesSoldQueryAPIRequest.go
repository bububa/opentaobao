package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradesSoldQueryAPIRequest
根据收件人信息查询交易单号 API请求
taobao.trades.sold.query

根据收件人信息查询交易单号。 */
type TaobaoTradesSoldQueryAPIRequest struct {
	model.Params
	// 查询条件列表，多个条件之间是OR关系，最多支持20个。receiver_name、receiver_mobile、receiver_phone至少有一个值不为空。
	_queryList []OrderQuery
}

// New
