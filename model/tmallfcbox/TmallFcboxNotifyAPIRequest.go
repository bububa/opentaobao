package tmallfcbox

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallfcboxnotifyAPIRequest 丰巢通知接口 API请求
// tmall.fcbox.notify
//
// tmax接收丰巢快递通知
type TmallfcboxnotifyAPIRequest struct {
	model.Params
	// 状态备注
	_stateRemark string
	// 申请记录当前状态
	_state string
	// 变更时间
	_operateTime string
	// 变更操作
	_operate string
	// 申请接口返回的申请标识
	_applyId string
}

// NewTmallfcboxnotifyRequest 初始化TmallfcboxnotifyAPIRequest对象
func NewTmallfcboxnotifyRequest() *TmallfcboxnotifyAPIRequest {
	return &TmallfcboxnotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallfcboxnotifyAPIRequest) GetApiMethodName() string {
	return "tmall.fcbox.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallfcboxnotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallfcboxnotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStateRemark is StateRemark Setter
// 状态备注
func (r *TmallfcboxnotifyAPIRequest) SetStateRemark(_stateRemark string) error {
	r._stateRemark = _stateRemark
	r.Set("state_remark", _stateRemark)
	return nil
}

// GetStateRemark StateRemark Getter
func (r TmallfcboxnotifyAPIRequest) GetStateRemark() string {
	return r._stateRemark
}

// SetState is State Setter
// 申请记录当前状态
func (r *TmallfcboxnotifyAPIRequest) SetState(_state string) error {
	r._state = _state
	r.Set("state", _state)
	return nil
}

// GetState State Getter
func (r TmallfcboxnotifyAPIRequest) GetState() string {
	return r._state
}

// SetOperateTime is OperateTime Setter
// 变更时间
func (r *TmallfcboxnotifyAPIRequest) SetOperateTime(_operateTime string) error {
	r._operateTime = _operateTime
	r.Set("operate_time", _operateTime)
	return nil
}

// GetOperateTime OperateTime Getter
func (r TmallfcboxnotifyAPIRequest) GetOperateTime() string {
	return r._operateTime
}

// SetOperate is Operate Setter
// 变更操作
func (r *TmallfcboxnotifyAPIRequest) SetOperate(_operate string) error {
	r._operate = _operate
	r.Set("operate", _operate)
	return nil
}

// GetOperate Operate Getter
func (r TmallfcboxnotifyAPIRequest) GetOperate() string {
	return r._operate
}

// SetApplyId is ApplyId Setter
// 申请接口返回的申请标识
func (r *TmallfcboxnotifyAPIRequest) SetApplyId(_applyId string) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TmallfcboxnotifyAPIRequest) GetApplyId() string {
	return r._applyId
}
