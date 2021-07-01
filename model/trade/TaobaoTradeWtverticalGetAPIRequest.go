package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeWtverticalGetAPIRequest
网厅垂直信息查询接口 API请求
taobao.trade.wtvertical.get

网厅订单垂直信息的查询 */
type TaobaoTradeWtverticalGetAPIRequest struct {
	model.Params
	// 主订单列表,用“，”分隔tid的字符串,最大列表长度为15
	_tids []int64
}

// New
