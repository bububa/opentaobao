package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
工单揽件 APIRequest
tmall.servicecenter.workcard.collect

服务工单揽件接口
*/
type TmallServicecenterWorkcardCollectRequest struct {
    model.Params

    // 工单id
    workcardId   int64 

    // 买家id
    buyerId   int64 

    // 扩展信息
    attributes   string 

}

func NewTmallServicecenterWorkcardCollectRequest() *TmallServicecenterWorkcardCollectRequest{
    return &TmallServicecenterWorkcardCollectRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkcardCollectRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.collect"
}

func (r TmallServicecenterWorkcardCollectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkcardCollectRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

func (r TmallServicecenterWorkcardCollectRequest) GetWorkcardId() int64 {
    return r.workcardId
}

func (r *TmallServicecenterWorkcardCollectRequest) SetBuyerId(buyerId int64) error {
    r.buyerId = buyerId
    r.Set("buyer_id", buyerId)
    return nil
}

func (r TmallServicecenterWorkcardCollectRequest) GetBuyerId() int64 {
    return r.buyerId
}

func (r *TmallServicecenterWorkcardCollectRequest) SetAttributes(attributes string) error {
    r.attributes = attributes
    r.Set("attributes", attributes)
    return nil
}

func (r TmallServicecenterWorkcardCollectRequest) GetAttributes() string {
    return r.attributes
}

