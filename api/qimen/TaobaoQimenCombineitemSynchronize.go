package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
组合商品接口 
taobao.qimen.combineitem.synchronize

ERP调用奇门的接口,将商品信息同步给WMS
*/
func TaobaoQimenCombineitemSynchronize(clt *core.SDKClient, req *qimen.TaobaoQimenCombineitemSynchronizeRequest, session string) (*qimen.TaobaoQimenCombineitemSynchronizeResponse, error) {
    var resp qimen.TaobaoQimenCombineitemSynchronizeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
