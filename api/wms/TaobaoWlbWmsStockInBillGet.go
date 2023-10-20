package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// TaobaoWlbWmsStockInBillGet 获取入库单信息
// taobao.wlb.wms.stock.in.bill.get
//
// 获取入库单信息
func TaobaoWlbWmsStockInBillGet(clt *core.SDKClient, req *wms.TaobaoWlbWmsStockInBillGetAPIRequest, session string) (*wms.TaobaoWlbWmsStockInBillGetAPIResponse, error) {
	var resp wms.TaobaoWlbWmsStockInBillGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
