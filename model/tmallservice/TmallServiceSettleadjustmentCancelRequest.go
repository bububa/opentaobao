package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消结算调整单 APIRequest
tmall.service.settleadjustment.cancel

提供给服务商在对取消已经发起的结算调整单。
通过说明调整单ID进行结算调整单取消。
*/
type TmallServiceSettleadjustmentCancelRequest struct {
    model.Params

    // 结算调整单ID
    id   int64 

    // 取消原因说明
    comments   string 

}

func NewTmallServiceSettleadjustmentCancelRequest() *TmallServiceSettleadjustmentCancelRequest{
    return &TmallServiceSettleadjustmentCancelRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServiceSettleadjustmentCancelRequest) GetApiMethodName() string {
    return "tmall.service.settleadjustment.cancel"
}

func (r TmallServiceSettleadjustmentCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServiceSettleadjustmentCancelRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r TmallServiceSettleadjustmentCancelRequest) GetId() int64 {
    return r.id
}

func (r *TmallServiceSettleadjustmentCancelRequest) SetComments(comments string) error {
    r.comments = comments
    r.Set("comments", comments)
    return nil
}

func (r TmallServiceSettleadjustmentCancelRequest) GetComments() string {
    return r.comments
}

