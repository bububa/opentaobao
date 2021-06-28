package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
商品同步接口 
taobao.qimen.singleitem.synchronize

ERP调用奇门的接口,同步商品信息给WMS
*/
func TaobaoQimenSingleitemSynchronize(clt *core.SDKClient, req *qimen.TaobaoQimenSingleitemSynchronizeRequest, session string) (*qimen.TaobaoQimenSingleitemSynchronizeAPIResponse, error) {
    var resp qimen.TaobaoQimenSingleitemSynchronizeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
