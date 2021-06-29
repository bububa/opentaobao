package tmallfcbox

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
丰巢通知接口 APIRequest
tmall.fcbox.notify

tmax接收丰巢快递通知
*/
type TmallFcboxNotifyRequest struct {
    model.Params

    // 申请接口返回的申请标识
    applyId   string 

    // 状态备注
    stateRemark   string 

    // 申请记录当前状态
    state   string 

    // 变更时间
    operateTime   string 

    // 变更操作
    operate   string 

}

func NewTmallFcboxNotifyRequest() *TmallFcboxNotifyRequest{
    return &TmallFcboxNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r TmallFcboxNotifyRequest) GetApiMethodName() string {
    return "tmall.fcbox.notify"
}

func (r TmallFcboxNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallFcboxNotifyRequest) SetApplyId(applyId string) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

func (r TmallFcboxNotifyRequest) GetApplyId() string {
    return r.applyId
}

func (r *TmallFcboxNotifyRequest) SetStateRemark(stateRemark string) error {
    r.stateRemark = stateRemark
    r.Set("state_remark", stateRemark)
    return nil
}

func (r TmallFcboxNotifyRequest) GetStateRemark() string {
    return r.stateRemark
}

func (r *TmallFcboxNotifyRequest) SetState(state string) error {
    r.state = state
    r.Set("state", state)
    return nil
}

func (r TmallFcboxNotifyRequest) GetState() string {
    return r.state
}

func (r *TmallFcboxNotifyRequest) SetOperateTime(operateTime string) error {
    r.operateTime = operateTime
    r.Set("operate_time", operateTime)
    return nil
}

func (r TmallFcboxNotifyRequest) GetOperateTime() string {
    return r.operateTime
}

func (r *TmallFcboxNotifyRequest) SetOperate(operate string) error {
    r.operate = operate
    r.Set("operate", operate)
    return nil
}

func (r TmallFcboxNotifyRequest) GetOperate() string {
    return r.operate
}

