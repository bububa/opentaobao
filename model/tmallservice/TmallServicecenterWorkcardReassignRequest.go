package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
工单改派门店 API请求
tmall.servicecenter.workcard.reassign

工单改派门店
*/
type TmallServicecenterWorkcardReassignRequest struct {
    model.Params
    // 请求入参
    reassignStoreRequest   *ReassignStoreRequest
}

// 初始化TmallServicecenterWorkcardReassignRequest对象
func NewTmallServicecenterWorkcardReassignRequest() *TmallServicecenterWorkcardReassignRequest{
    return &TmallServicecenterWorkcardReassignRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardReassignRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.reassign"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardReassignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReassignStoreRequest Setter
// 请求入参
func (r *TmallServicecenterWorkcardReassignRequest) SetReassignStoreRequest(reassignStoreRequest *ReassignStoreRequest) error {
    r.reassignStoreRequest = reassignStoreRequest
    r.Set("reassign_store_request", reassignStoreRequest)
    return nil
}

// ReassignStoreRequest Getter
func (r TmallServicecenterWorkcardReassignRequest) GetReassignStoreRequest() *ReassignStoreRequest {
    return r.reassignStoreRequest
}
