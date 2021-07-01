package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbTradeorderGetAPIRequest
根据交易号获取物流宝订单 API请求
taobao.wlb.tradeorder.get

根据交易类型和交易id查询物流宝订单详情 */
type TaobaoWlbTradeorderGetAPIRequest struct {
	model.Params
	// 子交易号
	_subTradeId string
	// 指定交易类型的交易号
	_tradeId string
	// 交易类型: TAOBAO--淘宝交易 OTHER_TRADE--其它交易
	_tradeType string
}

// New
