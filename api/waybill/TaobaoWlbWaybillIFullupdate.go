package waybill

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/waybill"
)

/* 
面单信息更新接口v1.0 
taobao.wlb.waybill.i.fullupdate

商家更新电子面单号对应的订单信息。<br/><br/>a.涉及到订单信息（如拆/合单、取消订单等）、修改发货地的逻辑时候，需要使用CANCEL+GET组合操作。一般的面单信息更新使用该接口；<br/><br/>b.该接口是全量修改接口，对于不做更新的字段也要把原有的字段值传进去，否则做为修改为空处理。
*/
func TaobaoWlbWaybillIFullupdate(clt *core.SDKClient, req *waybill.TaobaoWlbWaybillIFullupdateRequest, session string) (*waybill.TaobaoWlbWaybillIFullupdateAPIResponse, error) {
    var resp waybill.TaobaoWlbWaybillIFullupdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
