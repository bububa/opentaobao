package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophyrefundcsapplyrenderAPIRequest 商家代客售后逆向申请渲染获取 API请求
// alibaba.tcls.aelophy.refund.csapplyrender
//
// 提供商家代客售后逆向申请渲染获取的接口
type AlibabatclsaelophyrefundcsapplyrenderAPIRequest struct {
	model.Params
	// 系统自动生成
	_refundCsApplyRenderDTO *RefundCsApplyRenderDto
}

// NewAlibabatclsaelophyrefundcsapplyrenderRequest 初始化AlibabatclsaelophyrefundcsapplyrenderAPIRequest对象
func NewAlibabatclsaelophyrefundcsapplyrenderRequest() *AlibabatclsaelophyrefundcsapplyrenderAPIRequest {
	return &AlibabatclsaelophyrefundcsapplyrenderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophyrefundcsapplyrenderAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.refund.csapplyrender"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophyrefundcsapplyrenderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophyrefundcsapplyrenderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundCsApplyRenderDTO is RefundCsApplyRenderDTO Setter
// 系统自动生成
func (r *AlibabatclsaelophyrefundcsapplyrenderAPIRequest) SetRefundCsApplyRenderDTO(_refundCsApplyRenderDTO *RefundCsApplyRenderDto) error {
	r._refundCsApplyRenderDTO = _refundCsApplyRenderDTO
	r.Set("refund_cs_apply_render_d_t_o", _refundCsApplyRenderDTO)
	return nil
}

// GetRefundCsApplyRenderDTO RefundCsApplyRenderDTO Getter
func (r AlibabatclsaelophyrefundcsapplyrenderAPIRequest) GetRefundCsApplyRenderDTO() *RefundCsApplyRenderDto {
	return r._refundCsApplyRenderDTO
}
