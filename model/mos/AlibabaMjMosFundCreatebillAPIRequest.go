package mos

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMjMosFundCreatebillAPIRequest) Reset() {
	r._billDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjMosFundCreatebillAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.mos.fund.createbill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjMosFundCreatebillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMjMosFundCreatebillAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaMjMosFundCreatebillAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMjMosFundCreatebillRequest()
	},
}

// GetAlibabaMjMosFundCreatebillRequest 从 sync.Pool 获取 AlibabaMjMosFundCreatebillAPIRequest
func GetAlibabaMjMosFundCreatebillAPIRequest() *AlibabaMjMosFundCreatebillAPIRequest {
	return poolAlibabaMjMosFundCreatebillAPIRequest.Get().(*AlibabaMjMosFundCreatebillAPIRequest)
}

// ReleaseAlibabaMjMosFundCreatebillAPIRequest 将 AlibabaMjMosFundCreatebillAPIRequest 放入 sync.Pool
func ReleaseAlibabaMjMosFundCreatebillAPIRequest(v *AlibabaMjMosFundCreatebillAPIRequest) {
	v.Reset()
	poolAlibabaMjMosFundCreatebillAPIRequest.Put(v)
}
