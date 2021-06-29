package deliveryvoucher

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/deliveryvoucher"
)

/* 
预约接口 
taobao.game.deliveryvoucher.ordervoucher

提货券发券接口：同步券和订单的关联信息
*/
func TaobaoGameDeliveryvoucherOrdervoucher(clt *core.SDKClient, req *deliveryvoucher.TaobaoGameDeliveryvoucherOrdervoucherRequest, session string) (*deliveryvoucher.TaobaoGameDeliveryvoucherOrdervoucherAPIResponse, error) {
    var resp deliveryvoucher.TaobaoGameDeliveryvoucherOrdervoucherAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
