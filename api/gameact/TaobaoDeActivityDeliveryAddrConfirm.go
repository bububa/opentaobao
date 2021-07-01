package gameact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/gameact"
)

/* 
用户收件地址确认 
taobao.de.activity.delivery.addr.confirm

用户收件地址确认
*/
func TaobaoDeActivityDeliveryAddrConfirm(clt *core.SDKClient, req *gameact.TaobaoDeActivityDeliveryAddrConfirmAPIRequest, session string) (*gameact.TaobaoDeActivityDeliveryAddrConfirmAPIResponse, error) {
    var resp gameact.TaobaoDeActivityDeliveryAddrConfirmAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
