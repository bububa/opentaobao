package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjmosfundcancelbillAPIRequest 取消付款单 API请求
// alibaba.mj.mos.fund.cancelbill
//
// 取消付款单
type AlibabamjmosfundcancelbillAPIRequest struct {
	model.Params
	// 取消入参
	_cancelBillDTO *CancelBillDto
}

// NewAlibabamjmosfundcancelbillRequest 初始化AlibabamjmosfundcancelbillAPIRequest对象
func NewAlibabamjmosfundcancelbillRequest() *AlibabamjmosfundcancelbillAPIRequest {
	return &AlibabamjmosfundcancelbillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamjmosfundcancelbillAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.mos.fund.cancelbill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamjmosfundcancelbillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamjmosfundcancelbillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCancelBillDTO is CancelBillDTO Setter
// 取消入参
func (r *AlibabamjmosfundcancelbillAPIRequest) SetCancelBillDTO(_cancelBillDTO *CancelBillDto) error {
	r._cancelBillDTO = _cancelBillDTO
	r.Set("cancel_bill_d_t_o", _cancelBillDTO)
	return nil
}

// GetCancelBillDTO CancelBillDTO Getter
func (r AlibabamjmosfundcancelbillAPIRequest) GetCancelBillDTO() *CancelBillDto {
	return r._cancelBillDTO
}
