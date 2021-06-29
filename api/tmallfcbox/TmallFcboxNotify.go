package tmallfcbox

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallfcbox"
)

/* 
丰巢通知接口 
tmall.fcbox.notify

tmax接收丰巢快递通知
*/
func TmallFcboxNotify(clt *core.SDKClient, req *tmallfcbox.TmallFcboxNotifyRequest, session string) (*tmallfcbox.TmallFcboxNotifyAPIResponse, error) {
    var resp tmallfcbox.TmallFcboxNotifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
