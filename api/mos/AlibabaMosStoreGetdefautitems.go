package mos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mos"
)

/* 
获取默认状态下商品列表 
alibaba.mos.store.getdefautitems

获取默认状态下商品列表
*/
func AlibabaMosStoreGetdefautitems(clt *core.SDKClient, req *mos.AlibabaMosStoreGetdefautitemsAPIRequest, session string) (*mos.AlibabaMosStoreGetdefautitemsAPIResponse, error) {
    var resp mos.AlibabaMosStoreGetdefautitemsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
