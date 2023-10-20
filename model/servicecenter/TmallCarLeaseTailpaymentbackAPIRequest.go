package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseTailpaymentbackAPIRequest 尾款处置方案回传 API请求
// tmall.car.lease.tailpaymentback
//
// 尾款处置方案回传
type TmallCarLeaseTailpaymentbackAPIRequest struct {
	model.Params
	// 尾款方案
	_tailPaymentDTO *TailPaymentDto
}

// NewTmallCarLeaseTailpaymentbackRequest 初始化TmallCarLeaseTailpaymentbackAPIRequest对象
func NewTmallCarLeaseTailpaymentbackRequest() *TmallCarLeaseTailpaymentbackAPIRequest {
	return &TmallCarLeaseTailpaymentbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarLeaseTailpaymentbackAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.tailpaymentback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarLeaseTailpaymentbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarLeaseTailpaymentbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTailPaymentDTO is TailPaymentDTO Setter
// 尾款方案
func (r *TmallCarLeaseTailpaymentbackAPIRequest) SetTailPaymentDTO(_tailPaymentDTO *TailPaymentDto) error {
	r._tailPaymentDTO = _tailPaymentDTO
	r.Set("tail_payment_d_t_o", _tailPaymentDTO)
	return nil
}

// GetTailPaymentDTO TailPaymentDTO Getter
func (r TmallCarLeaseTailpaymentbackAPIRequest) GetTailPaymentDTO() *TailPaymentDto {
	return r._tailPaymentDTO
}
