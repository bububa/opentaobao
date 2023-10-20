package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmvouchersendAPIRequest 发送券给指定用户 API请求
// alibaba.alsc.crm.voucher.send
//
// 发送券给指定用户
type AlibabaalsccrmvouchersendAPIRequest struct {
	model.Params
	// 请求参数
	_paramVoucherSendOpenReq *VoucherSendOpenReq
}

// NewAlibabaalsccrmvouchersendRequest 初始化AlibabaalsccrmvouchersendAPIRequest对象
func NewAlibabaalsccrmvouchersendRequest() *AlibabaalsccrmvouchersendAPIRequest {
	return &AlibabaalsccrmvouchersendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmvouchersendAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.voucher.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmvouchersendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmvouchersendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamVoucherSendOpenReq is ParamVoucherSendOpenReq Setter
// 请求参数
func (r *AlibabaalsccrmvouchersendAPIRequest) SetParamVoucherSendOpenReq(_paramVoucherSendOpenReq *VoucherSendOpenReq) error {
	r._paramVoucherSendOpenReq = _paramVoucherSendOpenReq
	r.Set("param_voucher_send_open_req", _paramVoucherSendOpenReq)
	return nil
}

// GetParamVoucherSendOpenReq ParamVoucherSendOpenReq Getter
func (r AlibabaalsccrmvouchersendAPIRequest) GetParamVoucherSendOpenReq() *VoucherSendOpenReq {
	return r._paramVoucherSendOpenReq
}
