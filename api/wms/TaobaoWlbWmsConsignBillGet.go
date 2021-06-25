package wms

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wms"
)

/* 
获取销售订单发货信息 
taobao.wlb.wms.consign.bill.get

获取销售订单发货信息
*/
func TaobaoWlbWmsConsignBillGet(clt *core.SDKClient, req *wms.TaobaoWlbWmsConsignBillGetRequest, session string) (*wms.TaobaoWlbWmsConsignBillGetResponse, error) {
    var resp wms.TaobaoWlbWmsConsignBillGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
