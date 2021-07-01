package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
后端商品库存增量更新接口 
taobao.jst.astrolabe.storeinventory.update

增量更新门店或电商仓库存，该接口一次可以同时增量更新多个门店的多个商品的非确定性库存
*/
func TaobaoJstAstrolabeStoreinventoryUpdate(clt *core.SDKClient, req *omniorder.TaobaoJstAstrolabeStoreinventoryUpdateAPIRequest, session string) (*omniorder.TaobaoJstAstrolabeStoreinventoryUpdateAPIResponse, error) {
    var resp omniorder.TaobaoJstAstrolabeStoreinventoryUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
