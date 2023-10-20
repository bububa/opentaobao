package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophyrefundcsapplynewAPIRequest 代客退 API请求
// alibaba.tcls.aelophy.refund.csapply.new
//
// 代客退
type AlibabatclsaelophyrefundcsapplynewAPIRequest struct {
	model.Params
	// 逆向申请入参
	_refundCsApplyDTO *RefundCsApplyNewDto
}

// NewAlibabatclsaelophyrefundcsapplynewRequest 初始化AlibabatclsaelophyrefundcsapplynewAPIRequest对象
func NewAlibabatclsaelophyrefundcsapplynewRequest() *AlibabatclsaelophyrefundcsapplynewAPIRequest {
	return &AlibabatclsaelophyrefundcsapplynewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophyrefundcsapplynewAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.refund.csapply.new"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophyrefundcsapplynewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophyrefundcsapplynewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundCsApplyDTO is RefundCsApplyDTO Setter
// 逆向申请入参
func (r *AlibabatclsaelophyrefundcsapplynewAPIRequest) SetRefundCsApplyDTO(_refundCsApplyDTO *RefundCsApplyNewDto) error {
	r._refundCsApplyDTO = _refundCsApplyDTO
	r.Set("refund_cs_apply_d_t_o", _refundCsApplyDTO)
	return nil
}

// GetRefundCsApplyDTO RefundCsApplyDTO Getter
func (r AlibabatclsaelophyrefundcsapplynewAPIRequest) GetRefundCsApplyDTO() *RefundCsApplyNewDto {
	return r._refundCsApplyDTO
}
