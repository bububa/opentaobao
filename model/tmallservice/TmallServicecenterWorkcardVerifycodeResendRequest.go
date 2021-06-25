package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
重发核销码 APIRequest
tmall.servicecenter.workcard.verifycode.resend

重发核销码
*/
type TmallServicecenterWorkcardVerifycodeResendRequest struct {
    model.Params

    // 工单id
    workcardId   int64 

    // 门店/网点id
    serviceStoreId   int64 

}

func NewTmallServicecenterWorkcardVerifycodeResendRequest() *TmallServicecenterWorkcardVerifycodeResendRequest{
    return &TmallServicecenterWorkcardVerifycodeResendRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkcardVerifycodeResendRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.verifycode.resend"
}

func (r TmallServicecenterWorkcardVerifycodeResendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkcardVerifycodeResendRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

func (r TmallServicecenterWorkcardVerifycodeResendRequest) GetWorkcardId() int64 {
    return r.workcardId
}

func (r *TmallServicecenterWorkcardVerifycodeResendRequest) SetServiceStoreId(serviceStoreId int64) error {
    r.serviceStoreId = serviceStoreId
    r.Set("service_store_id", serviceStoreId)
    return nil
}

func (r TmallServicecenterWorkcardVerifycodeResendRequest) GetServiceStoreId() int64 {
    return r.serviceStoreId
}

