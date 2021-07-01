package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
门店信息查询接口 
taobao.qimen.store.query

商家在ERP等系统中调用该接口，查询门店相关信息
*/
func TaobaoQimenStoreQuery(clt *core.SDKClient, req *qimen.TaobaoQimenStoreQueryAPIRequest, session string) (*qimen.TaobaoQimenStoreQueryAPIResponse, error) {
    var resp qimen.TaobaoQimenStoreQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
