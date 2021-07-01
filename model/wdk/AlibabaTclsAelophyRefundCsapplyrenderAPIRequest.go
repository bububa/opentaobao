package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsAelophyRefundCsapplyrenderAPIRequest
商家代客售后逆向申请渲染获取 API请求
alibaba.tcls.aelophy.refund.csapplyrender

提供商家代客售后逆向申请渲染获取的接口 */
type AlibabaTclsAelophyRefundCsapplyrenderAPIRequest struct {
	model.Params
	// 系统自动生成
	_refundCsApplyRenderDTO *RefundCsApplyRenderDto
}

// NewAlibabaTclsAelophyRefundCsapplyrenderRequest 初始化AlibabaTclsAelophyRefundCsapplyrenderAPIRequest对象
func NewAlibabaTclsAelophyRefundCsapplyrenderRequest() *AlibabaTclsAelophyRefundCsapplyrenderAPIRequest {
	return &AlibabaTclsAelophyRefundCsapplyrenderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyRefundCsapplyrenderAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.refund.csapplyrender"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyRefundCsapplyrenderAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefundCsApplyRenderDTO Setter
// 系统自动生成
func (r *AlibabaTclsAelophyRefundCsapplyrenderAPIRequest) SetRefundCsApplyRenderDTO(_refundCsApplyRenderDTO *RefundCsApplyRenderDto) error {
	r._refundCsApplyRenderDTO = _refundCsApplyRenderDTO
	r.Set("refund_cs_apply_render_d_t_o", _refundCsApplyRenderDTO)
	return nil
}

// Get RefundCsApplyRenderDTO Getter
func (r AlibabaTclsAelophyRefundCsapplyrenderAPIRequest) GetRefundCsApplyRenderDTO() *RefundCsApplyRenderDto {
	return r._refundCsApplyRenderDTO
}
