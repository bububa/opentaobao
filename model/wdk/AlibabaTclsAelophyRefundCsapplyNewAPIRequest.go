package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyRefundCsapplyNewAPIRequest 代客退 API请求
// alibaba.tcls.aelophy.refund.csapply.new
//
// 代客退
type AlibabaTclsAelophyRefundCsapplyNewAPIRequest struct {
	model.Params
	// 逆向申请入参
	_refundCsApplyDTO *RefundCsApplyNewDto
}

// NewAlibabaTclsAelophyRefundCsapplyNewRequest 初始化AlibabaTclsAelophyRefundCsapplyNewAPIRequest对象
func NewAlibabaTclsAelophyRefundCsapplyNewRequest() *AlibabaTclsAelophyRefundCsapplyNewAPIRequest {
	return &AlibabaTclsAelophyRefundCsapplyNewAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTclsAelophyRefundCsapplyNewAPIRequest) Reset() {
	r._refundCsApplyDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyRefundCsapplyNewAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.refund.csapply.new"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyRefundCsapplyNewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyRefundCsapplyNewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundCsApplyDTO is RefundCsApplyDTO Setter
// 逆向申请入参
func (r *AlibabaTclsAelophyRefundCsapplyNewAPIRequest) SetRefundCsApplyDTO(_refundCsApplyDTO *RefundCsApplyNewDto) error {
	r._refundCsApplyDTO = _refundCsApplyDTO
	r.Set("refund_cs_apply_d_t_o", _refundCsApplyDTO)
	return nil
}

// GetRefundCsApplyDTO RefundCsApplyDTO Getter
func (r AlibabaTclsAelophyRefundCsapplyNewAPIRequest) GetRefundCsApplyDTO() *RefundCsApplyNewDto {
	return r._refundCsApplyDTO
}

var poolAlibabaTclsAelophyRefundCsapplyNewAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTclsAelophyRefundCsapplyNewRequest()
	},
}

// GetAlibabaTclsAelophyRefundCsapplyNewRequest 从 sync.Pool 获取 AlibabaTclsAelophyRefundCsapplyNewAPIRequest
func GetAlibabaTclsAelophyRefundCsapplyNewAPIRequest() *AlibabaTclsAelophyRefundCsapplyNewAPIRequest {
	return poolAlibabaTclsAelophyRefundCsapplyNewAPIRequest.Get().(*AlibabaTclsAelophyRefundCsapplyNewAPIRequest)
}

// ReleaseAlibabaTclsAelophyRefundCsapplyNewAPIRequest 将 AlibabaTclsAelophyRefundCsapplyNewAPIRequest 放入 sync.Pool
func ReleaseAlibabaTclsAelophyRefundCsapplyNewAPIRequest(v *AlibabaTclsAelophyRefundCsapplyNewAPIRequest) {
	v.Reset()
	poolAlibabaTclsAelophyRefundCsapplyNewAPIRequest.Put(v)
}
