package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
门店新增接口 
taobao.qimen.store.create

isv调用接口来讲线下门店同步到线上
*/
func TaobaoQimenStoreCreate(clt *core.SDKClient, req *qimen.TaobaoQimenStoreCreateAPIRequest, session string) (*qimen.TaobaoQimenStoreCreateAPIResponse, error) {
    var resp qimen.TaobaoQimenStoreCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
