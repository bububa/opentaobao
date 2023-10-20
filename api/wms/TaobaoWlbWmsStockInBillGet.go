package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// TaobaoWlbWmsStockInBillGet 获取入库单信息
// taobao.wlb.wms.stock.in.bill.get
//
// 获取入库单信息
func TaobaoWlbWmsStockInBillGet(clt *core.SDKClient, req *wms.TaobaoWlbWmsStockInBillGetAPIRequest, resp *wms.TaobaoWlbWmsStockInBillGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
