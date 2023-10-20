package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// TaobaoWlbWmsStockOutBillGet 通过订单号获取单个出库单发货信息
// taobao.wlb.wms.stock.out.bill.get
//
// 通过订单号获取单个出库单发货信息
func TaobaoWlbWmsStockOutBillGet(clt *core.SDKClient, req *wms.TaobaoWlbWmsStockOutBillGetAPIRequest, resp *wms.TaobaoWlbWmsStockOutBillGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
