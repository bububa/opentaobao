package lsttrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lsttrade"
)

// AlibabaLstTradeShiporderQuery 供应商数据开放--发货单接口
// alibaba.lst.trade.shiporder.query
//
// 供应商数据开放--发货单接口
func AlibabaLstTradeShiporderQuery(clt *core.SDKClient, req *lsttrade.AlibabaLstTradeShiporderQueryAPIRequest, resp *lsttrade.AlibabaLstTradeShiporderQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
