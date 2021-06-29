package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商确认服务完成 API请求
tmall.servicecenter.workcard.confirm

提供给外部合作服务商，用于通知天猫，告知寄修服务厂内操作全部完成
*/
type TmallServicecenterWorkcardConfirmRequest struct {
    model.Params
    // 工单id
    workcardId   int64
}

// 初始化TmallServicecenterWorkcardConfirmRequest对象
func NewTmallServicecenterWorkcardConfirmRequest() *TmallServicecenterWorkcardConfirmRequest{
    return &TmallServicecenterWorkcardConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardConfirmRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardConfirmRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

// WorkcardId Getter
func (r TmallServicecenterWorkcardConfirmRequest) GetWorkcardId() int64 {
    return r.workcardId
}
