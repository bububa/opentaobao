package nrt

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nrt"
)

/* 
同步摊位收银比例 
tmall.nrt.stall.payratio.synchronize

ISV同步摊位收银比例到阿里
*/
func TmallNrtStallPayratioSynchronize(clt *core.SDKClient, req *nrt.TmallNrtStallPayratioSynchronizeAPIRequest, session string) (*nrt.TmallNrtStallPayratioSynchronizeAPIResponse, error) {
    var resp nrt.TmallNrtStallPayratioSynchronizeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
