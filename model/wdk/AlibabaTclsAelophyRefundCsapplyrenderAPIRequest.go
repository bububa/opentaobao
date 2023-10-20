package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyRefundCsapplyrenderAPIRequest 商家代客售后逆向申请渲染获取 API请求
// alibaba.tcls.aelophy.refund.csapplyrender
//
// 提供商家代客售后逆向申请渲染获取的接口
type AlibabaTclsAelophyRefundCsapplyrenderAPIRequest struct {
	model.Params
	// 系统自动生成
	_refundCsApplyRenderDTO *RefundCsApplyRenderDto
}

// NewAlibabaTclsAelophyRefundCsapplyrenderRequest 初始化AlibabaTclsAelophyRefundCsapplyrenderAPIRequest对象
func NewAlibabaTclsAelophyRefundCsapplyrenderRequest() *AlibabaTclsAelophyRefundCsapplyrenderAPIRequest {
	return &AlibabaTclsAelophyRefundCsapplyrenderAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTclsAelophyRefundCsapplyrenderAPIRequest) Reset() {
	r._refundCsApplyRenderDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyRefundCsapplyrenderAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.refund.csapplyrender"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyRefundCsapplyrenderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyRefundCsapplyrenderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundCsApplyRenderDTO is RefundCsApplyRenderDTO Setter
// 系统自动生成
func (r *AlibabaTclsAelophyRefundCsapplyrenderAPIRequest) SetRefundCsApplyRenderDTO(_refundCsApplyRenderDTO *RefundCsApplyRenderDto) error {
	r._refundCsApplyRenderDTO = _refundCsApplyRenderDTO
	r.Set("refund_cs_apply_render_d_t_o", _refundCsApplyRenderDTO)
	return nil
}

// GetRefundCsApplyRenderDTO RefundCsApplyRenderDTO Getter
func (r AlibabaTclsAelophyRefundCsapplyrenderAPIRequest) GetRefundCsApplyRenderDTO() *RefundCsApplyRenderDto {
	return r._refundCsApplyRenderDTO
}

var poolAlibabaTclsAelophyRefundCsapplyrenderAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTclsAelophyRefundCsapplyrenderRequest()
	},
}

// GetAlibabaTclsAelophyRefundCsapplyrenderRequest 从 sync.Pool 获取 AlibabaTclsAelophyRefundCsapplyrenderAPIRequest
func GetAlibabaTclsAelophyRefundCsapplyrenderAPIRequest() *AlibabaTclsAelophyRefundCsapplyrenderAPIRequest {
	return poolAlibabaTclsAelophyRefundCsapplyrenderAPIRequest.Get().(*AlibabaTclsAelophyRefundCsapplyrenderAPIRequest)
}

// ReleaseAlibabaTclsAelophyRefundCsapplyrenderAPIRequest 将 AlibabaTclsAelophyRefundCsapplyrenderAPIRequest 放入 sync.Pool
func ReleaseAlibabaTclsAelophyRefundCsapplyrenderAPIRequest(v *AlibabaTclsAelophyRefundCsapplyrenderAPIRequest) {
	v.Reset()
	poolAlibabaTclsAelophyRefundCsapplyrenderAPIRequest.Put(v)
}
