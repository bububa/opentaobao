package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商创建核销单 APIRequest
tmall.servicecenter.identifytask.create

服务商调用该接口进行创建核销单操作
*/
type TmallServicecenterIdentifytaskCreateRequest struct {
    model.Params

    // 工单列表
    workcardIds   []int64 

    // 是否改派
    reassign   bool 

    // 服务商自定义的外部核销单id
    outerId   string 

}

func NewTmallServicecenterIdentifytaskCreateRequest() *TmallServicecenterIdentifytaskCreateRequest{
    return &TmallServicecenterIdentifytaskCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterIdentifytaskCreateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.identifytask.create"
}

func (r TmallServicecenterIdentifytaskCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterIdentifytaskCreateRequest) SetWorkcardIds(workcardIds []int64) error {
    r.workcardIds = workcardIds
    r.Set("workcard_ids", workcardIds)
    return nil
}

func (r TmallServicecenterIdentifytaskCreateRequest) GetWorkcardIds() []int64 {
    return r.workcardIds
}

func (r *TmallServicecenterIdentifytaskCreateRequest) SetReassign(reassign bool) error {
    r.reassign = reassign
    r.Set("reassign", reassign)
    return nil
}

func (r TmallServicecenterIdentifytaskCreateRequest) GetReassign() bool {
    return r.reassign
}

func (r *TmallServicecenterIdentifytaskCreateRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TmallServicecenterIdentifytaskCreateRequest) GetOuterId() string {
    return r.outerId
}

