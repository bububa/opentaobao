package tmallfcbox

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
丰巢通知接口 API请求
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

// 初始化TmallFcboxNotifyRequest对象
func NewTmallFcboxNotifyRequest() *TmallFcboxNotifyRequest{
    return &TmallFcboxNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallFcboxNotifyRequest) GetApiMethodName() string {
    return "tmall.fcbox.notify"
}

// IRequest interface 方法, 获取API参数
func (r TmallFcboxNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApplyId Setter
// 申请接口返回的申请标识
func (r *TmallFcboxNotifyRequest) SetApplyId(applyId string) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

// ApplyId Getter
func (r TmallFcboxNotifyRequest) GetApplyId() string {
    return r.applyId
}
// StateRemark Setter
// 状态备注
func (r *TmallFcboxNotifyRequest) SetStateRemark(stateRemark string) error {
    r.stateRemark = stateRemark
    r.Set("state_remark", stateRemark)
    return nil
}

// StateRemark Getter
func (r TmallFcboxNotifyRequest) GetStateRemark() string {
    return r.stateRemark
}
// State Setter
// 申请记录当前状态
func (r *TmallFcboxNotifyRequest) SetState(state string) error {
    r.state = state
    r.Set("state", state)
    return nil
}

// State Getter
func (r TmallFcboxNotifyRequest) GetState() string {
    return r.state
}
// OperateTime Setter
// 变更时间
func (r *TmallFcboxNotifyRequest) SetOperateTime(operateTime string) error {
    r.operateTime = operateTime
    r.Set("operate_time", operateTime)
    return nil
}

// OperateTime Getter
func (r TmallFcboxNotifyRequest) GetOperateTime() string {
    return r.operateTime
}
// Operate Setter
// 变更操作
func (r *TmallFcboxNotifyRequest) SetOperate(operate string) error {
    r.operate = operate
    r.Set("operate", operate)
    return nil
}

// Operate Getter
func (r TmallFcboxNotifyRequest) GetOperate() string {
    return r.operate
}
