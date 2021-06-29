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
    _applyId   string
    // 状态备注
    _stateRemark   string
    // 申请记录当前状态
    _state   string
    // 变更时间
    _operateTime   string
    // 变更操作
    _operate   string
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
func (r *TmallFcboxNotifyRequest) SetApplyId(_applyId string) error {
    r._applyId = _applyId
    r.Set("apply_id", _applyId)
    return nil
}

// ApplyId Getter
func (r TmallFcboxNotifyRequest) GetApplyId() string {
    return r._applyId
}
// StateRemark Setter
// 状态备注
func (r *TmallFcboxNotifyRequest) SetStateRemark(_stateRemark string) error {
    r._stateRemark = _stateRemark
    r.Set("state_remark", _stateRemark)
    return nil
}

// StateRemark Getter
func (r TmallFcboxNotifyRequest) GetStateRemark() string {
    return r._stateRemark
}
// State Setter
// 申请记录当前状态
func (r *TmallFcboxNotifyRequest) SetState(_state string) error {
    r._state = _state
    r.Set("state", _state)
    return nil
}

// State Getter
func (r TmallFcboxNotifyRequest) GetState() string {
    return r._state
}
// OperateTime Setter
// 变更时间
func (r *TmallFcboxNotifyRequest) SetOperateTime(_operateTime string) error {
    r._operateTime = _operateTime
    r.Set("operate_time", _operateTime)
    return nil
}

// OperateTime Getter
func (r TmallFcboxNotifyRequest) GetOperateTime() string {
    return r._operateTime
}
// Operate Setter
// 变更操作
func (r *TmallFcboxNotifyRequest) SetOperate(_operate string) error {
    r._operate = _operate
    r.Set("operate", _operate)
    return nil
}

// Operate Getter
func (r TmallFcboxNotifyRequest) GetOperate() string {
    return r._operate
}
