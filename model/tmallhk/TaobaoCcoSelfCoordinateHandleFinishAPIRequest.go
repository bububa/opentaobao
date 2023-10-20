package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCcoSelfCoordinateHandleFinishAPIRequest 天猫国际直购供应商处理完结回复通知 API请求
// taobao.cco.self.coordinate.handle.finish
//
// 天猫国际直购供应商处理完结回复通知
type TaobaoCcoSelfCoordinateHandleFinishAPIRequest struct {
	model.Params
	// 处理人
	_operator string
	// 回复模板,其中key/value均为字符串
	_replyData string
	// 工单编号
	_caseId int64
}

// NewTaobaoCcoSelfCoordinateHandleFinishRequest 初始化TaobaoCcoSelfCoordinateHandleFinishAPIRequest对象
func NewTaobaoCcoSelfCoordinateHandleFinishRequest() *TaobaoCcoSelfCoordinateHandleFinishAPIRequest {
	return &TaobaoCcoSelfCoordinateHandleFinishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCcoSelfCoordinateHandleFinishAPIRequest) GetApiMethodName() string {
	return "taobao.cco.self.coordinate.handle.finish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCcoSelfCoordinateHandleFinishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCcoSelfCoordinateHandleFinishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOperator is Operator Setter
// 处理人
func (r *TaobaoCcoSelfCoordinateHandleFinishAPIRequest) SetOperator(_operator string) error {
	r._operator = _operator
	r.Set("operator", _operator)
	return nil
}

// GetOperator Operator Getter
func (r TaobaoCcoSelfCoordinateHandleFinishAPIRequest) GetOperator() string {
	return r._operator
}

// SetReplyData is ReplyData Setter
// 回复模板,其中key/value均为字符串
func (r *TaobaoCcoSelfCoordinateHandleFinishAPIRequest) SetReplyData(_replyData string) error {
	r._replyData = _replyData
	r.Set("reply_data", _replyData)
	return nil
}

// GetReplyData ReplyData Getter
func (r TaobaoCcoSelfCoordinateHandleFinishAPIRequest) GetReplyData() string {
	return r._replyData
}

// SetCaseId is CaseId Setter
// 工单编号
func (r *TaobaoCcoSelfCoordinateHandleFinishAPIRequest) SetCaseId(_caseId int64) error {
	r._caseId = _caseId
	r.Set("case_id", _caseId)
	return nil
}

// GetCaseId CaseId Getter
func (r TaobaoCcoSelfCoordinateHandleFinishAPIRequest) GetCaseId() int64 {
	return r._caseId
}
