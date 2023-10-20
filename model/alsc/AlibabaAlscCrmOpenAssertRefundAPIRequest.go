package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmOpenAssertRefundAPIRequest 资产核销回退接口 API请求
// alibaba.alsc.crm.open.assert.refund
//
// 回退已经核销的储值积分券资产
type AlibabaAlscCrmOpenAssertRefundAPIRequest struct {
	model.Params
	// 入参
	_paramPropertyRefundOpenReq *PropertyRefundOpenReq
}

// NewAlibabaAlscCrmOpenAssertRefundRequest 初始化AlibabaAlscCrmOpenAssertRefundAPIRequest对象
func NewAlibabaAlscCrmOpenAssertRefundRequest() *AlibabaAlscCrmOpenAssertRefundAPIRequest {
	return &AlibabaAlscCrmOpenAssertRefundAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmOpenAssertRefundAPIRequest) Reset() {
	r._paramPropertyRefundOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmOpenAssertRefundAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.open.assert.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmOpenAssertRefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmOpenAssertRefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPropertyRefundOpenReq is ParamPropertyRefundOpenReq Setter
// 入参
func (r *AlibabaAlscCrmOpenAssertRefundAPIRequest) SetParamPropertyRefundOpenReq(_paramPropertyRefundOpenReq *PropertyRefundOpenReq) error {
	r._paramPropertyRefundOpenReq = _paramPropertyRefundOpenReq
	r.Set("param_property_refund_open_req", _paramPropertyRefundOpenReq)
	return nil
}

// GetParamPropertyRefundOpenReq ParamPropertyRefundOpenReq Getter
func (r AlibabaAlscCrmOpenAssertRefundAPIRequest) GetParamPropertyRefundOpenReq() *PropertyRefundOpenReq {
	return r._paramPropertyRefundOpenReq
}

var poolAlibabaAlscCrmOpenAssertRefundAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmOpenAssertRefundRequest()
	},
}

// GetAlibabaAlscCrmOpenAssertRefundRequest 从 sync.Pool 获取 AlibabaAlscCrmOpenAssertRefundAPIRequest
func GetAlibabaAlscCrmOpenAssertRefundAPIRequest() *AlibabaAlscCrmOpenAssertRefundAPIRequest {
	return poolAlibabaAlscCrmOpenAssertRefundAPIRequest.Get().(*AlibabaAlscCrmOpenAssertRefundAPIRequest)
}

// ReleaseAlibabaAlscCrmOpenAssertRefundAPIRequest 将 AlibabaAlscCrmOpenAssertRefundAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmOpenAssertRefundAPIRequest(v *AlibabaAlscCrmOpenAssertRefundAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmOpenAssertRefundAPIRequest.Put(v)
}
