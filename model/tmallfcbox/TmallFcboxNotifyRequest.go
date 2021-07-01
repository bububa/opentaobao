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
type TmallFcboxNotifyAPIRequest struct {
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

// 初始化TmallFcboxNotifyAPIRequest对象
func NewTmallFcboxNotifyRequest() *TmallFcboxNotifyAPIRequest{
    return &TmallFcboxNotifyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallFcboxNotifyAPIRequest) GetApiMethodName() string {
    return "tmall.fcbox.notify"
}

// IRequest interface 方法, 获取API参数
func (r TmallFcboxNotifyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApplyId Setter
// 申请接口返回的申请标识
func (r *TmallFcboxNotifyAPIRequest) SetApplyId(_applyId string) error {
    r._applyId = _applyId
    r.Set("apply_id", _applyId)
    return nil
}

// ApplyId Getter
func (r TmallFcboxNotifyAPIRequest) GetApplyId() string {
    return r._applyId
}
// StateRemark Setter
// 状态备注
func (r *TmallFcboxNotifyAPIRequest) SetStateRemark(_stateRemark string) error {
    r._stateRemark = _stateRemark
    r.Set("state_remark", _stateRemark)
    return nil
}

// StateRemark Getter
func (r TmallFcboxNotifyAPIRequest) GetStateRemark() string {
    return r._stateRemark
}
// State Setter
// 申请记录当前状态
func (r *TmallFcboxNotifyAPIRequest) SetState(_state string) error {
    r._state = _state
    r.Set("state", _state)
    return nil
}

// State Getter
func (r TmallFcboxNotifyAPIRequest) GetState() string {
    return r._state
}
// OperateTime Setter
// 变更时间
func (r *TmallFcboxNotifyAPIRequest) SetOperateTime(_operateTime string) error {
    r._operateTime = _operateTime
    r.Set("operate_time", _operateTime)
    return nil
}

// OperateTime Getter
func (r TmallFcboxNotifyAPIRequest) GetOperateTime() string {
    return r._operateTime
}
// Operate Setter
// 变更操作
func (r *TmallFcboxNotifyAPIRequest) SetOperate(_operate string) error {
    r._operate = _operate
    r.Set("operate", _operate)
    return nil
}

// Operate Getter
func (r TmallFcboxNotifyAPIRequest) GetOperate() string {
    return r._operate
}
