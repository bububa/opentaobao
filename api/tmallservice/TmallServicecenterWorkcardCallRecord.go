package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
回访记录回传API 
tmall.servicecenter.workcard.call.record

客满回访信息登记
*/
func TmallServicecenterWorkcardCallRecord(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardCallRecordRequest, session string) (*tmallservice.TmallServicecenterWorkcardCallRecordResponse, error) {
    var resp tmallservice.TmallServicecenterWorkcardCallRecordAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
