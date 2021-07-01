package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
分单结果同步给星盘 
taobao.omniorder.allocatedinfo.sync

ISV分单完成，将分单结果同步给星盘
*/
func TaobaoOmniorderAllocatedinfoSync(clt *core.SDKClient, req *omniorder.TaobaoOmniorderAllocatedinfoSyncAPIRequest, session string) (*omniorder.TaobaoOmniorderAllocatedinfoSyncAPIResponse, error) {
    var resp omniorder.TaobaoOmniorderAllocatedinfoSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
