package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmOpenAssertVerifyAPIRequest 资产核销接口 API请求
// alibaba.alsc.crm.open.assert.verify
//
// 核销储值，积分，券资产
type AlibabaAlscCrmOpenAssertVerifyAPIRequest struct {
	model.Params
	// 入参
	_paramPropertyVerifyOpenReq *PropertyVerifyOpenReq
}

// NewAlibabaAlscCrmOpenAssertVerifyRequest 初始化AlibabaAlscCrmOpenAssertVerifyAPIRequest对象
func NewAlibabaAlscCrmOpenAssertVerifyRequest() *AlibabaAlscCrmOpenAssertVerifyAPIRequest {
	return &AlibabaAlscCrmOpenAssertVerifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmOpenAssertVerifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.open.assert.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmOpenAssertVerifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamPropertyVerifyOpenReq is ParamPropertyVerifyOpenReq Setter
// 入参
func (r *AlibabaAlscCrmOpenAssertVerifyAPIRequest) SetParamPropertyVerifyOpenReq(_paramPropertyVerifyOpenReq *PropertyVerifyOpenReq) error {
	r._paramPropertyVerifyOpenReq = _paramPropertyVerifyOpenReq
	r.Set("param_property_verify_open_req", _paramPropertyVerifyOpenReq)
	return nil
}

// GetParamPropertyVerifyOpenReq ParamPropertyVerifyOpenReq Getter
func (r AlibabaAlscCrmOpenAssertVerifyAPIRequest) GetParamPropertyVerifyOpenReq() *PropertyVerifyOpenReq {
	return r._paramPropertyVerifyOpenReq
}
