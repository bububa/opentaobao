package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmvoucherstatuschangeAPIRequest 优惠券状态更改 API请求
// alibaba.alsc.crm.voucher.status.change
//
// 核销优惠券
type AlibabaalsccrmvoucherstatuschangeAPIRequest struct {
	model.Params
	// 参数
	_paramVoucherStatusChangeOpenReq *VoucherStatusChangeOpenReq
}

// NewAlibabaalsccrmvoucherstatuschangeRequest 初始化AlibabaalsccrmvoucherstatuschangeAPIRequest对象
func NewAlibabaalsccrmvoucherstatuschangeRequest() *AlibabaalsccrmvoucherstatuschangeAPIRequest {
	return &AlibabaalsccrmvoucherstatuschangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmvoucherstatuschangeAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.voucher.status.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmvoucherstatuschangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmvoucherstatuschangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamVoucherStatusChangeOpenReq is ParamVoucherStatusChangeOpenReq Setter
// 参数
func (r *AlibabaalsccrmvoucherstatuschangeAPIRequest) SetParamVoucherStatusChangeOpenReq(_paramVoucherStatusChangeOpenReq *VoucherStatusChangeOpenReq) error {
	r._paramVoucherStatusChangeOpenReq = _paramVoucherStatusChangeOpenReq
	r.Set("param_voucher_status_change_open_req", _paramVoucherStatusChangeOpenReq)
	return nil
}

// GetParamVoucherStatusChangeOpenReq ParamVoucherStatusChangeOpenReq Getter
func (r AlibabaalsccrmvoucherstatuschangeAPIRequest) GetParamVoucherStatusChangeOpenReq() *VoucherStatusChangeOpenReq {
	return r._paramVoucherStatusChangeOpenReq
}
