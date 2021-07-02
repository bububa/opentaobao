package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjMosFundCancelbillAPIRequest 取消付款单 API请求
// alibaba.mj.mos.fund.cancelbill
//
// 取消付款单
type AlibabaMjMosFundCancelbillAPIRequest struct {
	model.Params
	// 取消入参
	_cancelBillDTO *CancelBillDto
}

// NewAlibabaMjMosFundCancelbillRequest 初始化AlibabaMjMosFundCancelbillAPIRequest对象
func NewAlibabaMjMosFundCancelbillRequest() *AlibabaMjMosFundCancelbillAPIRequest {
	return &AlibabaMjMosFundCancelbillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjMosFundCancelbillAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.mos.fund.cancelbill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjMosFundCancelbillAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCancelBillDTO is CancelBillDTO Setter
// 取消入参
func (r *AlibabaMjMosFundCancelbillAPIRequest) SetCancelBillDTO(_cancelBillDTO *CancelBillDto) error {
	r._cancelBillDTO = _cancelBillDTO
	r.Set("cancel_bill_d_t_o", _cancelBillDTO)
	return nil
}

// GetCancelBillDTO CancelBillDTO Getter
func (r AlibabaMjMosFundCancelbillAPIRequest) GetCancelBillDTO() *CancelBillDto {
	return r._cancelBillDTO
}
