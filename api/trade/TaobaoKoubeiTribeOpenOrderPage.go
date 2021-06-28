package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
口碑综合体订单列表信息查询 
taobao.koubei.tribe.open.order.page

查询口碑商圈用户的订单列表信息
*/
func TaobaoKoubeiTribeOpenOrderPage(clt *core.SDKClient, req *trade.TaobaoKoubeiTribeOpenOrderPageRequest, session string) (*trade.TaobaoKoubeiTribeOpenOrderPageAPIResponse, error) {
    var resp trade.TaobaoKoubeiTribeOpenOrderPageAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
