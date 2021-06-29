package deliveryvoucher

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/deliveryvoucher"
)

/* 
提货券发货接口 
taobao.game.deliveryvoucher.sendgoods

提货券发券接口：同步券和订单的关联信息
*/
func TaobaoGameDeliveryvoucherSendgoods(clt *core.SDKClient, req *deliveryvoucher.TaobaoGameDeliveryvoucherSendgoodsRequest, session string) (*deliveryvoucher.TaobaoGameDeliveryvoucherSendgoodsAPIResponse, error) {
    var resp deliveryvoucher.TaobaoGameDeliveryvoucherSendgoodsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
