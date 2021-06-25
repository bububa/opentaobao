package wms

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wms"
)

/* 
通过订单号获取单个出库单发货信息 
taobao.wlb.wms.stock.out.bill.get

通过订单号获取单个出库单发货信息
*/
func TaobaoWlbWmsStockOutBillGet(clt *core.SDKClient, req *wms.TaobaoWlbWmsStockOutBillGetRequest, session string) (*wms.TaobaoWlbWmsStockOutBillGetResponse, error) {
    var resp wms.TaobaoWlbWmsStockOutBillGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
