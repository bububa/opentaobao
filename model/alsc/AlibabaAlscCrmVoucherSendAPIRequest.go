package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmVoucherSendAPIRequest
发送券给指定用户 API请求
alibaba.alsc.crm.voucher.send

发送券给指定用户 */
type AlibabaAlscCrmVoucherSendAPIRequest struct {
	model.Params
	// 请求参数
	_paramVoucherSendOpenReq *VoucherSendOpenReq
}

// NewAlibabaAlscCrmVoucherSendRequest 初始化AlibabaAlscCrmVoucherSendAPIRequest对象
func NewAlibabaAlscCrmVoucherSendRequest() *AlibabaAlscCrmVoucherSendAPIRequest {
	return &AlibabaAlscCrmVoucherSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmVoucherSendAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.voucher.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmVoucherSendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamVoucherSendOpenReq Setter
// 请求参数
func (r *AlibabaAlscCrmVoucherSendAPIRequest) SetParamVoucherSendOpenReq(_paramVoucherSendOpenReq *VoucherSendOpenReq) error {
	r._paramVoucherSendOpenReq = _paramVoucherSendOpenReq
	r.Set("param_voucher_send_open_req", _paramVoucherSendOpenReq)
	return nil
}

// Get ParamVoucherSendOpenReq Getter
func (r AlibabaAlscCrmVoucherSendAPIRequest) GetParamVoucherSendOpenReq() *VoucherSendOpenReq {
	return r._paramVoucherSendOpenReq
}
