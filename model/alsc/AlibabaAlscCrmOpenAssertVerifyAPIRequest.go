package alsc

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmOpenAssertVerifyAPIRequest) Reset() {
	r._paramPropertyVerifyOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmOpenAssertVerifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.open.assert.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmOpenAssertVerifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmOpenAssertVerifyAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlscCrmOpenAssertVerifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmOpenAssertVerifyRequest()
	},
}

// GetAlibabaAlscCrmOpenAssertVerifyRequest 从 sync.Pool 获取 AlibabaAlscCrmOpenAssertVerifyAPIRequest
func GetAlibabaAlscCrmOpenAssertVerifyAPIRequest() *AlibabaAlscCrmOpenAssertVerifyAPIRequest {
	return poolAlibabaAlscCrmOpenAssertVerifyAPIRequest.Get().(*AlibabaAlscCrmOpenAssertVerifyAPIRequest)
}

// ReleaseAlibabaAlscCrmOpenAssertVerifyAPIRequest 将 AlibabaAlscCrmOpenAssertVerifyAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmOpenAssertVerifyAPIRequest(v *AlibabaAlscCrmOpenAssertVerifyAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmOpenAssertVerifyAPIRequest.Put(v)
}
