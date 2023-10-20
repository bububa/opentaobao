package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardUpdatelogisticsAPIRequest 更新物流进度 API请求
// tmall.servicecenter.workcard.updatelogistics
//
// 提供给外部合作服务商的物流进度更改接口
type TmallServicecenterWorkcardUpdatelogisticsAPIRequest struct {
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

// NewTmallServicecenterWorkcardUpdatelogisticsRequest 初始化TmallServicecenterWorkcardUpdatelogisticsAPIRequest对象
func NewTmallServicecenterWorkcardUpdatelogisticsRequest() *TmallServicecenterWorkcardUpdatelogisticsAPIRequest {
	return &TmallServicecenterWorkcardUpdatelogisticsAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterWorkcardUpdatelogisticsAPIRequest) Reset() {
	r._action = ""
	r._expressCompany = ""
	r._expressCode = ""
	r._workcardId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardUpdatelogisticsAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.updatelogistics"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardUpdatelogisticsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkcardUpdatelogisticsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAction is Action Setter
// 工单操作
func (r *TmallServicecenterWorkcardUpdatelogisticsAPIRequest) SetAction(_action string) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// GetAction Action Getter
func (r TmallServicecenterWorkcardUpdatelogisticsAPIRequest) GetAction() string {
	return r._action
}

// SetExpressCompany is ExpressCompany Setter
// 快递公司
func (r *TmallServicecenterWorkcardUpdatelogisticsAPIRequest) SetExpressCompany(_expressCompany string) error {
	r._expressCompany = _expressCompany
	r.Set("express_company", _expressCompany)
	return nil
}

// GetExpressCompany ExpressCompany Getter
func (r TmallServicecenterWorkcardUpdatelogisticsAPIRequest) GetExpressCompany() string {
	return r._expressCompany
}

// SetExpressCode is ExpressCode Setter
// 快递号
func (r *TmallServicecenterWorkcardUpdatelogisticsAPIRequest) SetExpressCode(_expressCode string) error {
	r._expressCode = _expressCode
	r.Set("express_code", _expressCode)
	return nil
}

// GetExpressCode ExpressCode Getter
func (r TmallServicecenterWorkcardUpdatelogisticsAPIRequest) GetExpressCode() string {
	return r._expressCode
}

// SetWorkcardId is WorkcardId Setter
// 工单号
func (r *TmallServicecenterWorkcardUpdatelogisticsAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallServicecenterWorkcardUpdatelogisticsAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

var poolTmallServicecenterWorkcardUpdatelogisticsAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterWorkcardUpdatelogisticsRequest()
	},
}

// GetTmallServicecenterWorkcardUpdatelogisticsRequest 从 sync.Pool 获取 TmallServicecenterWorkcardUpdatelogisticsAPIRequest
func GetTmallServicecenterWorkcardUpdatelogisticsAPIRequest() *TmallServicecenterWorkcardUpdatelogisticsAPIRequest {
	return poolTmallServicecenterWorkcardUpdatelogisticsAPIRequest.Get().(*TmallServicecenterWorkcardUpdatelogisticsAPIRequest)
}

// ReleaseTmallServicecenterWorkcardUpdatelogisticsAPIRequest 将 TmallServicecenterWorkcardUpdatelogisticsAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterWorkcardUpdatelogisticsAPIRequest(v *TmallServicecenterWorkcardUpdatelogisticsAPIRequest) {
	v.Reset()
	poolTmallServicecenterWorkcardUpdatelogisticsAPIRequest.Put(v)
}
