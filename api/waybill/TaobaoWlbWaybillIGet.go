package waybill

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/waybill"
)

/* 
获取物流服务商电子面单号v1.0 
taobao.wlb.waybill.i.get

商家根据订单信息，实时、批量获取指定物流服务商的电子面单号。
*/
func TaobaoWlbWaybillIGet(clt *core.SDKClient, req *waybill.TaobaoWlbWaybillIGetRequest, session string) (*waybill.TaobaoWlbWaybillIGetResponse, error) {
    var resp waybill.TaobaoWlbWaybillIGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
