package uscesl

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/uscesl"
)

/* 
新增电子价签商家门店接口 
taobao.uscesl.biz.store.insert

新增电子价签商家门店接口
*/
func TaobaoUsceslBizStoreInsert(clt *core.SDKClient, req *uscesl.TaobaoUsceslBizStoreInsertRequest, session string) (*uscesl.TaobaoUsceslBizStoreInsertAPIResponse, error) {
    var resp uscesl.TaobaoUsceslBizStoreInsertAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
