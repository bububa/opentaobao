package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// CainiaoBimTradeorderConsign 驱动保税交易订单发货
// cainiao.bim.tradeorder.consign
//
// 驱动保税交易订单发货
func CainiaoBimTradeorderConsign(clt *core.SDKClient, req *wms.CainiaoBimTradeorderConsignAPIRequest, resp *wms.CainiaoBimTradeorderConsignAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
