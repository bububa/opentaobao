package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjMosFundCreatebillAPIRequest 创建一个付款单 API请求
// alibaba.mj.mos.fund.createbill
//
// 创建一个付款单
type AlibabaMjMosFundCreatebillAPIRequest struct {
	model.Params
	// 创建付款单入参
	_billDto *CreateBillDto
}

// NewAlibabaMjMosFundCreatebillRequest 初始化AlibabaMjMosFundCreatebillAPIRequest对象
func NewAlibabaMjMosFundCreatebillRequest() *AlibabaMjMosFundCreatebillAPIRequest {
	return &AlibabaMjMosFundCreatebillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjMosFundCreatebillAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.mos.fund.createbill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjMosFundCreatebillAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBillDto is BillDto Setter
// 创建付款单入参
func (r *AlibabaMjMosFundCreatebillAPIRequest) SetBillDto(_billDto *CreateBillDto) error {
	r._billDto = _billDto
	r.Set("bill_dto", _billDto)
	return nil
}

// GetBillDto BillDto Getter
func (r AlibabaMjMosFundCreatebillAPIRequest) GetBillDto() *CreateBillDto {
	return r._billDto
}
