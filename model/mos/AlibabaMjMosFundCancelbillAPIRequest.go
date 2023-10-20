package mos

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMjMosFundCancelbillAPIRequest) Reset() {
	r._cancelBillDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjMosFundCancelbillAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.mos.fund.cancelbill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjMosFundCancelbillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMjMosFundCancelbillAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaMjMosFundCancelbillAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMjMosFundCancelbillRequest()
	},
}

// GetAlibabaMjMosFundCancelbillRequest 从 sync.Pool 获取 AlibabaMjMosFundCancelbillAPIRequest
func GetAlibabaMjMosFundCancelbillAPIRequest() *AlibabaMjMosFundCancelbillAPIRequest {
	return poolAlibabaMjMosFundCancelbillAPIRequest.Get().(*AlibabaMjMosFundCancelbillAPIRequest)
}

// ReleaseAlibabaMjMosFundCancelbillAPIRequest 将 AlibabaMjMosFundCancelbillAPIRequest 放入 sync.Pool
func ReleaseAlibabaMjMosFundCancelbillAPIRequest(v *AlibabaMjMosFundCancelbillAPIRequest) {
	v.Reset()
	poolAlibabaMjMosFundCancelbillAPIRequest.Put(v)
}
