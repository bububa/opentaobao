package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophyrefundcsapplyAPIRequest 商家代客售后提交逆向申请 API请求
// alibaba.tcls.aelophy.refund.csapply
//
// 商家代客售后提交逆向申请
type AlibabatclsaelophyrefundcsapplyAPIRequest struct {
	model.Params
	// 逆向申请入参
	_refundCsApplyDTO *RefundCsApplyDto
}

// NewAlibabatclsaelophyrefundcsapplyRequest 初始化AlibabatclsaelophyrefundcsapplyAPIRequest对象
func NewAlibabatclsaelophyrefundcsapplyRequest() *AlibabatclsaelophyrefundcsapplyAPIRequest {
	return &AlibabatclsaelophyrefundcsapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophyrefundcsapplyAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.refund.csapply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophyrefundcsapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophyrefundcsapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundCsApplyDTO is RefundCsApplyDTO Setter
// 逆向申请入参
func (r *AlibabatclsaelophyrefundcsapplyAPIRequest) SetRefundCsApplyDTO(_refundCsApplyDTO *RefundCsApplyDto) error {
	r._refundCsApplyDTO = _refundCsApplyDTO
	r.Set("refund_cs_apply_d_t_o", _refundCsApplyDTO)
	return nil
}

// GetRefundCsApplyDTO RefundCsApplyDTO Getter
func (r AlibabatclsaelophyrefundcsapplyAPIRequest) GetRefundCsApplyDTO() *RefundCsApplyDto {
	return r._refundCsApplyDTO
}
