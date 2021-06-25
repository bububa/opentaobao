package servicecenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/servicecenter"
)

/* 
签收接口 
tmall.msf.receive

签收接口
*/
func TmallMsfReceive(clt *core.SDKClient, req *servicecenter.TmallMsfReceiveRequest, session string) (*servicecenter.TmallMsfReceiveResponse, error) {
    var resp servicecenter.TmallMsfReceiveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
