package waybill

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/waybill"
)

/* 
商家取消获取的电子面单号v1.0 
taobao.wlb.waybill.i.cancel

面单号有误需要取消的时候，调用该接口取消获取的电子面单。
*/
func TaobaoWlbWaybillICancel(clt *core.SDKClient, req *waybill.TaobaoWlbWaybillICancelAPIRequest, session string) (*waybill.TaobaoWlbWaybillICancelAPIResponse, error) {
    var resp waybill.TaobaoWlbWaybillICancelAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
