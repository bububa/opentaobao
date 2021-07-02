package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyRefundCsapplyAPIRequest 商家代客售后提交逆向申请 API请求
// alibaba.tcls.aelophy.refund.csapply
//
// 商家代客售后提交逆向申请
type AlibabaTclsAelophyRefundCsapplyAPIRequest struct {
	model.Params
	// 逆向申请入参
	_refundCsApplyDTO *RefundCsApplyDto
}

// NewAlibabaTclsAelophyRefundCsapplyRequest 初始化AlibabaTclsAelophyRefundCsapplyAPIRequest对象
func NewAlibabaTclsAelophyRefundCsapplyRequest() *AlibabaTclsAelophyRefundCsapplyAPIRequest {
	return &AlibabaTclsAelophyRefundCsapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyRefundCsapplyAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.refund.csapply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyRefundCsapplyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefundCsApplyDTO Setter
// 逆向申请入参
func (r *AlibabaTclsAelophyRefundCsapplyAPIRequest) SetRefundCsApplyDTO(_refundCsApplyDTO *RefundCsApplyDto) error {
	r._refundCsApplyDTO = _refundCsApplyDTO
	r.Set("refund_cs_apply_d_t_o", _refundCsApplyDTO)
	return nil
}

// Get RefundCsApplyDTO Getter
func (r AlibabaTclsAelophyRefundCsapplyAPIRequest) GetRefundCsApplyDTO() *RefundCsApplyDto {
	return r._refundCsApplyDTO
}
