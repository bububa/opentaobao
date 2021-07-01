package lsttrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstTradeOrderGetAPIRequest
零售通交易订单查询--品牌商视角 API请求
alibaba.lst.trade.order.get

根据订单id查询零售通交易订单 */
type AlibabaLstTradeOrderGetAPIRequest struct {
	model.Params
	// 主订单id
	_mainOrderId int64
	// 子订单id
	_subOrderId int64
}

// New
