package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
门店删除接口 
taobao.qimen.store.delete

商家在ERP等系统中调用该接口，删除线下门店
*/
func TaobaoQimenStoreDelete(clt *core.SDKClient, req *qimen.TaobaoQimenStoreDeleteRequest, session string) (*qimen.TaobaoQimenStoreDeleteAPIResponse, error) {
    var resp qimen.TaobaoQimenStoreDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
