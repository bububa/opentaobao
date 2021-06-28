package tbk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tbk"
)

/* 
淘宝客-推广者-店铺搜索 
taobao.tbk.shop.get

淘宝客店铺查询
*/
func TaobaoTbkShopGet(clt *core.SDKClient, req *tbk.TaobaoTbkShopGetRequest, session string) (*tbk.TaobaoTbkShopGetAPIResponse, error) {
    var resp tbk.TaobaoTbkShopGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
