package tmallnr

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallnr"
)

/* 
ISV查询门店 
alibaba.lsy.crm.activity.store.getstorelist

ISV查询门店
*/
func AlibabaLsyCrmActivityStoreGetstorelist(clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmActivityStoreGetstorelistRequest, session string) (*tmallnr.AlibabaLsyCrmActivityStoreGetstorelistAPIResponse, error) {
    var resp tmallnr.AlibabaLsyCrmActivityStoreGetstorelistAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
