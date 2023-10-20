package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardConfirmAPIRequest 服务商确认服务完成 API请求
// tmall.servicecenter.workcard.confirm
//
// 提供给外部合作服务商，用于通知天猫，告知寄修服务厂内操作全部完成
type TmallServicecenterWorkcardConfirmAPIRequest struct {
	model.Params
	// 工单id
	_workcardId int64
}

// NewTmallServicecenterWorkcardConfirmRequest 初始化TmallServicecenterWorkcardConfirmAPIRequest对象
func NewTmallServicecenterWorkcardConfirmRequest() *TmallServicecenterWorkcardConfirmAPIRequest {
	return &TmallServicecenterWorkcardConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardConfirmAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkcardConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardConfirmAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallServicecenterWorkcardConfirmAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}
