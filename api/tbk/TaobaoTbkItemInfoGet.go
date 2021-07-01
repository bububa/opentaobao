package tbk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tbk"
)

/* 
淘宝客-公用-淘宝客商品详情查询(简版) 
taobao.tbk.item.info.get

淘宝客商品详情查询（简版）
*/
func TaobaoTbkItemInfoGet(clt *core.SDKClient, req *tbk.TaobaoTbkItemInfoGetAPIRequest, session string) (*tbk.TaobaoTbkItemInfoGetAPIResponse, error) {
    var resp tbk.TaobaoTbkItemInfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
