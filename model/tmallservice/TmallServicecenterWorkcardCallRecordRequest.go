package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
回访记录回传API APIRequest
tmall.servicecenter.workcard.call.record

客满回访信息登记
*/
type TmallServicecenterWorkcardCallRecordRequest struct {
    model.Params

    // 请求入参
    busiRequest   *UpdateAttributeRequest 

}

func NewTmallServicecenterWorkcardCallRecordRequest() *TmallServicecenterWorkcardCallRecordRequest{
    return &TmallServicecenterWorkcardCallRecordRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkcardCallRecordRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.call.record"
}

func (r TmallServicecenterWorkcardCallRecordRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkcardCallRecordRequest) SetBusiRequest(busiRequest *UpdateAttributeRequest) error {
    r.busiRequest = busiRequest
    r.Set("busi_request", busiRequest)
    return nil
}

func (r TmallServicecenterWorkcardCallRecordRequest) GetBusiRequest() *UpdateAttributeRequest {
    return r.busiRequest
}

