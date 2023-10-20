package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoInventoryAdjustExternal 非交易库存调整单
// taobao.inventory.adjust.external
//
// 商家非交易调整库存，调拨出库、盘点等时调用
func TaobaoInventoryAdjustExternal(clt *core.SDKClient, req *fenxiao.TaobaoInventoryAdjustExternalAPIRequest, resp *fenxiao.TaobaoInventoryAdjustExternalAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
