package lsttrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstTradeOrderQuerychangeAPIRequest
订单id批量查询（品牌商视角） API请求
alibaba.lst.trade.order.querychange

根据品牌和时间段查询有变更记录的订单id */
type AlibabaLstTradeOrderQuerychangeAPIRequest struct {
	model.Params
	// 入参包装类
	_query *LstOrderQuery
}

// New
