package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务工单拉取 APIRequest
tmall.servicecenter.task.get

接口供服务供应商通过交易主订单查询其未拉取的任务类工单
*/
type TmallServicecenterTaskGetRequest struct {
    model.Params

    // Taobao主交易订单ID
    parentBizOrderId   int64 

}

func NewTmallServicecenterTaskGetRequest() *TmallServicecenterTaskGetRequest{
    return &TmallServicecenterTaskGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterTaskGetRequest) GetApiMethodName() string {
    return "tmall.servicecenter.task.get"
}

func (r TmallServicecenterTaskGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterTaskGetRequest) SetParentBizOrderId(parentBizOrderId int64) error {
    r.parentBizOrderId = parentBizOrderId
    r.Set("parent_biz_order_id", parentBizOrderId)
    return nil
}

func (r TmallServicecenterTaskGetRequest) GetParentBizOrderId() int64 {
    return r.parentBizOrderId
}

