package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmVoucherSendAPIRequest 发送券给指定用户 API请求
// alibaba.alsc.crm.voucher.send
//
// 发送券给指定用户
type AlibabaAlscCrmVoucherSendAPIRequest struct {
	model.Params
	// 请求参数
	_paramVoucherSendOpenReq *VoucherSendOpenReq
}

// NewAlibabaAlscCrmVoucherSendRequest 初始化AlibabaAlscCrmVoucherSendAPIRequest对象
func NewAlibabaAlscCrmVoucherSendRequest() *AlibabaAlscCrmVoucherSendAPIRequest {
	return &AlibabaAlscCrmVoucherSendAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmVoucherSendAPIRequest) Reset() {
	r._paramVoucherSendOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmVoucherSendAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.voucher.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmVoucherSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmVoucherSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamVoucherSendOpenReq is ParamVoucherSendOpenReq Setter
// 请求参数
func (r *AlibabaAlscCrmVoucherSendAPIRequest) SetParamVoucherSendOpenReq(_paramVoucherSendOpenReq *VoucherSendOpenReq) error {
	r._paramVoucherSendOpenReq = _paramVoucherSendOpenReq
	r.Set("param_voucher_send_open_req", _paramVoucherSendOpenReq)
	return nil
}

// GetParamVoucherSendOpenReq ParamVoucherSendOpenReq Getter
func (r AlibabaAlscCrmVoucherSendAPIRequest) GetParamVoucherSendOpenReq() *VoucherSendOpenReq {
	return r._paramVoucherSendOpenReq
}

var poolAlibabaAlscCrmVoucherSendAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmVoucherSendRequest()
	},
}

// GetAlibabaAlscCrmVoucherSendRequest 从 sync.Pool 获取 AlibabaAlscCrmVoucherSendAPIRequest
func GetAlibabaAlscCrmVoucherSendAPIRequest() *AlibabaAlscCrmVoucherSendAPIRequest {
	return poolAlibabaAlscCrmVoucherSendAPIRequest.Get().(*AlibabaAlscCrmVoucherSendAPIRequest)
}

// ReleaseAlibabaAlscCrmVoucherSendAPIRequest 将 AlibabaAlscCrmVoucherSendAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmVoucherSendAPIRequest(v *AlibabaAlscCrmVoucherSendAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmVoucherSendAPIRequest.Put(v)
}
