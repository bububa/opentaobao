package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
店铺同步接口 
taobao.qimen.shop.synchronize

店铺同步接口描述
*/
func TaobaoQimenShopSynchronize(clt *core.SDKClient, req *qimen.TaobaoQimenShopSynchronizeRequest, session string) (*qimen.TaobaoQimenShopSynchronizeAPIResponse, error) {
    var resp qimen.TaobaoQimenShopSynchronizeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
