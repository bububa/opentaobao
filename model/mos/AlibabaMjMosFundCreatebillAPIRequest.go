package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjmosfundcreatebillAPIRequest 创建一个付款单 API请求
// alibaba.mj.mos.fund.createbill
//
// 创建一个付款单
type AlibabamjmosfundcreatebillAPIRequest struct {
	model.Params
	// 创建付款单入参
	_billDto *CreateBillDto
}

// NewAlibabamjmosfundcreatebillRequest 初始化AlibabamjmosfundcreatebillAPIRequest对象
func NewAlibabamjmosfundcreatebillRequest() *AlibabamjmosfundcreatebillAPIRequest {
	return &AlibabamjmosfundcreatebillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamjmosfundcreatebillAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.mos.fund.createbill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamjmosfundcreatebillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamjmosfundcreatebillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBillDto is BillDto Setter
// 创建付款单入参
func (r *AlibabamjmosfundcreatebillAPIRequest) SetBillDto(_billDto *CreateBillDto) error {
	r._billDto = _billDto
	r.Set("bill_dto", _billDto)
	return nil
}

// GetBillDto BillDto Getter
func (r AlibabamjmosfundcreatebillAPIRequest) GetBillDto() *CreateBillDto {
	return r._billDto
}
