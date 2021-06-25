package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
工单改派门店 APIRequest
tmall.servicecenter.workcard.reassign

工单改派门店
*/
type TmallServicecenterWorkcardReassignRequest struct {
    model.Params

    // 请求入参
    reassignStoreRequest   *ReassignStoreRequest 

}

func NewTmallServicecenterWorkcardReassignRequest() *TmallServicecenterWorkcardReassignRequest{
    return &TmallServicecenterWorkcardReassignRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkcardReassignRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.reassign"
}

func (r TmallServicecenterWorkcardReassignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkcardReassignRequest) SetReassignStoreRequest(reassignStoreRequest *ReassignStoreRequest) error {
    r.reassignStoreRequest = reassignStoreRequest
    r.Set("reassign_store_request", reassignStoreRequest)
    return nil
}

func (r TmallServicecenterWorkcardReassignRequest) GetReassignStoreRequest() *ReassignStoreRequest {
    return r.reassignStoreRequest
}

