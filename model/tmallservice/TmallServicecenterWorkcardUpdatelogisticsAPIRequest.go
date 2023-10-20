package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardupdatelogisticsAPIRequest 更新物流进度 API请求
// tmall.servicecenter.workcard.updatelogistics
//
// 提供给外部合作服务商的物流进度更改接口
type TmallservicecenterworkcardupdatelogisticsAPIRequest struct {
	model.Params
	// 工单操作
	_action string
	// 快递公司
	_expressCompany string
	// 快递号
	_expressCode string
	// 工单号
	_workcardId int64
}

// NewTmallservicecenterworkcardupdatelogisticsRequest 初始化TmallservicecenterworkcardupdatelogisticsAPIRequest对象
func NewTmallservicecenterworkcardupdatelogisticsRequest() *TmallservicecenterworkcardupdatelogisticsAPIRequest {
	return &TmallservicecenterworkcardupdatelogisticsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkcardupdatelogisticsAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.updatelogistics"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkcardupdatelogisticsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkcardupdatelogisticsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAction is Action Setter
// 工单操作
func (r *TmallservicecenterworkcardupdatelogisticsAPIRequest) SetAction(_action string) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// GetAction Action Getter
func (r TmallservicecenterworkcardupdatelogisticsAPIRequest) GetAction() string {
	return r._action
}

// SetExpressCompany is ExpressCompany Setter
// 快递公司
func (r *TmallservicecenterworkcardupdatelogisticsAPIRequest) SetExpressCompany(_expressCompany string) error {
	r._expressCompany = _expressCompany
	r.Set("express_company", _expressCompany)
	return nil
}

// GetExpressCompany ExpressCompany Getter
func (r TmallservicecenterworkcardupdatelogisticsAPIRequest) GetExpressCompany() string {
	return r._expressCompany
}

// SetExpressCode is ExpressCode Setter
// 快递号
func (r *TmallservicecenterworkcardupdatelogisticsAPIRequest) SetExpressCode(_expressCode string) error {
	r._expressCode = _expressCode
	r.Set("express_code", _expressCode)
	return nil
}

// GetExpressCode ExpressCode Getter
func (r TmallservicecenterworkcardupdatelogisticsAPIRequest) GetExpressCode() string {
	return r._expressCode
}

// SetWorkcardId is WorkcardId Setter
// 工单号
func (r *TmallservicecenterworkcardupdatelogisticsAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallservicecenterworkcardupdatelogisticsAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}
