package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTclsAelophyRefundCsapplyAPIRequest) Reset() {
	r._refundCsApplyDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyRefundCsapplyAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.refund.csapply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyRefundCsapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyRefundCsapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundCsApplyDTO is RefundCsApplyDTO Setter
// 逆向申请入参
func (r *AlibabaTclsAelophyRefundCsapplyAPIRequest) SetRefundCsApplyDTO(_refundCsApplyDTO *RefundCsApplyDto) error {
	r._refundCsApplyDTO = _refundCsApplyDTO
	r.Set("refund_cs_apply_d_t_o", _refundCsApplyDTO)
	return nil
}

// GetRefundCsApplyDTO RefundCsApplyDTO Getter
func (r AlibabaTclsAelophyRefundCsapplyAPIRequest) GetRefundCsApplyDTO() *RefundCsApplyDto {
	return r._refundCsApplyDTO
}

var poolAlibabaTclsAelophyRefundCsapplyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTclsAelophyRefundCsapplyRequest()
	},
}

// GetAlibabaTclsAelophyRefundCsapplyRequest 从 sync.Pool 获取 AlibabaTclsAelophyRefundCsapplyAPIRequest
func GetAlibabaTclsAelophyRefundCsapplyAPIRequest() *AlibabaTclsAelophyRefundCsapplyAPIRequest {
	return poolAlibabaTclsAelophyRefundCsapplyAPIRequest.Get().(*AlibabaTclsAelophyRefundCsapplyAPIRequest)
}

// ReleaseAlibabaTclsAelophyRefundCsapplyAPIRequest 将 AlibabaTclsAelophyRefundCsapplyAPIRequest 放入 sync.Pool
func ReleaseAlibabaTclsAelophyRefundCsapplyAPIRequest(v *AlibabaTclsAelophyRefundCsapplyAPIRequest) {
	v.Reset()
	poolAlibabaTclsAelophyRefundCsapplyAPIRequest.Put(v)
}
