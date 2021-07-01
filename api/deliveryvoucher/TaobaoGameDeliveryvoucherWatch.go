package deliveryvoucher

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/deliveryvoucher"
)

/* 
监控预约数据 
taobao.game.deliveryvoucher.watch

监控预约数据
*/
func TaobaoGameDeliveryvoucherWatch(clt *core.SDKClient, req *deliveryvoucher.TaobaoGameDeliveryvoucherWatchAPIRequest, session string) (*deliveryvoucher.TaobaoGameDeliveryvoucherWatchAPIResponse, error) {
    var resp deliveryvoucher.TaobaoGameDeliveryvoucherWatchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
