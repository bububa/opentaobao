package lsttrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstTradeShiporderQueryAPIRequest
供应商数据开放--发货单接口 API请求
alibaba.lst.trade.shiporder.query

供应商数据开放--发货单接口 */
type AlibabaLstTradeShiporderQueryAPIRequest struct {
	model.Params
	// 入参
	_paramLstShipOrderQuery *LstShipOrderQuery
}

// New
