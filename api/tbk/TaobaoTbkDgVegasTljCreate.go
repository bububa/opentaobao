package tbk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tbk"
)

/* 
淘宝客-推广者-淘礼金创建 
taobao.tbk.dg.vegas.tlj.create

创建淘礼金
*/
func TaobaoTbkDgVegasTljCreate(clt *core.SDKClient, req *tbk.TaobaoTbkDgVegasTljCreateAPIRequest, session string) (*tbk.TaobaoTbkDgVegasTljCreateAPIResponse, error) {
    var resp tbk.TaobaoTbkDgVegasTljCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
