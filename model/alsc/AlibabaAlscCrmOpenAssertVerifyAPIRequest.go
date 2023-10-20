package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmopenassertverifyAPIRequest 资产核销接口 API请求
// alibaba.alsc.crm.open.assert.verify
//
// 核销储值，积分，券资产
type AlibabaalsccrmopenassertverifyAPIRequest struct {
	model.Params
	// 入参
	_paramPropertyVerifyOpenReq *PropertyVerifyOpenReq
}

// NewAlibabaalsccrmopenassertverifyRequest 初始化AlibabaalsccrmopenassertverifyAPIRequest对象
func NewAlibabaalsccrmopenassertverifyRequest() *AlibabaalsccrmopenassertverifyAPIRequest {
	return &AlibabaalsccrmopenassertverifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmopenassertverifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.open.assert.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmopenassertverifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmopenassertverifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPropertyVerifyOpenReq is ParamPropertyVerifyOpenReq Setter
// 入参
func (r *AlibabaalsccrmopenassertverifyAPIRequest) SetParamPropertyVerifyOpenReq(_paramPropertyVerifyOpenReq *PropertyVerifyOpenReq) error {
	r._paramPropertyVerifyOpenReq = _paramPropertyVerifyOpenReq
	r.Set("param_property_verify_open_req", _paramPropertyVerifyOpenReq)
	return nil
}

// GetParamPropertyVerifyOpenReq ParamPropertyVerifyOpenReq Getter
func (r AlibabaalsccrmopenassertverifyAPIRequest) GetParamPropertyVerifyOpenReq() *PropertyVerifyOpenReq {
	return r._paramPropertyVerifyOpenReq
}
