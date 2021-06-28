package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
对一笔交易添加备注 
taobao.trade.memo.add

根据登录用户的身份（买家或卖家），自动添加相应的交易备注,不能重复调用些接口添加备注，需要更新备注请用taobao.trade.memo.update
*/
func TaobaoTradeMemoAdd(clt *core.SDKClient, req *trade.TaobaoTradeMemoAddRequest, session string) (*trade.TaobaoTradeMemoAddAPIResponse, error) {
    var resp trade.TaobaoTradeMemoAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
