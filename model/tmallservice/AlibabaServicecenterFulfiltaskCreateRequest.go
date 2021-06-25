package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
合单生成核销单 APIRequest
alibaba.servicecenter.fulfiltask.create

服务对工单进行合单，合单的结果是生成核销单
*/
type AlibabaServicecenterFulfiltaskCreateRequest struct {
    model.Params

    // 工单id列表
    workcardIds   []Number 

    // 外部单号
    outerId   string 

}

func NewAlibabaServicecenterFulfiltaskCreateRequest() *AlibabaServicecenterFulfiltaskCreateRequest{
    return &AlibabaServicecenterFulfiltaskCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaServicecenterFulfiltaskCreateRequest) GetApiMethodName() string {
    return "alibaba.servicecenter.fulfiltask.create"
}

func (r AlibabaServicecenterFulfiltaskCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaServicecenterFulfiltaskCreateRequest) SetWorkcardIds(workcardIds []Number) error {
    r.workcardIds = workcardIds
    r.Set("workcard_ids", workcardIds)
    return nil
}

func (r AlibabaServicecenterFulfiltaskCreateRequest) GetWorkcardIds() []Number {
    return r.workcardIds
}

func (r *AlibabaServicecenterFulfiltaskCreateRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r AlibabaServicecenterFulfiltaskCreateRequest) GetOuterId() string {
    return r.outerId
}

