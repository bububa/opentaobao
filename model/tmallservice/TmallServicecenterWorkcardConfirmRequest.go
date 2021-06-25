package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商确认服务完成 APIRequest
tmall.servicecenter.workcard.confirm

提供给外部合作服务商，用于通知天猫，告知寄修服务厂内操作全部完成
*/
type TmallServicecenterWorkcardConfirmRequest struct {
    model.Params

    // 工单id
    workcardId   int64 

}

func NewTmallServicecenterWorkcardConfirmRequest() *TmallServicecenterWorkcardConfirmRequest{
    return &TmallServicecenterWorkcardConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkcardConfirmRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.confirm"
}

func (r TmallServicecenterWorkcardConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkcardConfirmRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

func (r TmallServicecenterWorkcardConfirmRequest) GetWorkcardId() int64 {
    return r.workcardId
}

