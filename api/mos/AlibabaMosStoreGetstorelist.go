package mos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mos"
)

/* 
根据屏编号获取专柜集 
alibaba.mos.store.getstorelist

根据屏编号获取专柜集
*/
func AlibabaMosStoreGetstorelist(clt *core.SDKClient, req *mos.AlibabaMosStoreGetstorelistAPIRequest, session string) (*mos.AlibabaMosStoreGetstorelistAPIResponse, error) {
    var resp mos.AlibabaMosStoreGetstorelistAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
