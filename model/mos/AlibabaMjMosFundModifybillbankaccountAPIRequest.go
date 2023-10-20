package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjMosFundModifybillbankaccountAPIRequest 修改付款单的银行账户信息 API请求
// alibaba.mj.mos.fund.modifybillbankaccount
//
// 修改付款单的银行账户信息
type AlibabaMjMosFundModifybillbankaccountAPIRequest struct {
	model.Params
	// 修改入参
	_modifyDto *ModifyBillDto
}

// NewAlibabaMjMosFundModifybillbankaccountRequest 初始化AlibabaMjMosFundModifybillbankaccountAPIRequest对象
func NewAlibabaMjMosFundModifybillbankaccountRequest() *AlibabaMjMosFundModifybillbankaccountAPIRequest {
	return &AlibabaMjMosFundModifybillbankaccountAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMjMosFundModifybillbankaccountAPIRequest) Reset() {
	r._modifyDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjMosFundModifybillbankaccountAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.mos.fund.modifybillbankaccount"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjMosFundModifybillbankaccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMjMosFundModifybillbankaccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModifyDto is ModifyDto Setter
// 修改入参
func (r *AlibabaMjMosFundModifybillbankaccountAPIRequest) SetModifyDto(_modifyDto *ModifyBillDto) error {
	r._modifyDto = _modifyDto
	r.Set("modify_dto", _modifyDto)
	return nil
}

// GetModifyDto ModifyDto Getter
func (r AlibabaMjMosFundModifybillbankaccountAPIRequest) GetModifyDto() *ModifyBillDto {
	return r._modifyDto
}

var poolAlibabaMjMosFundModifybillbankaccountAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMjMosFundModifybillbankaccountRequest()
	},
}

// GetAlibabaMjMosFundModifybillbankaccountRequest 从 sync.Pool 获取 AlibabaMjMosFundModifybillbankaccountAPIRequest
func GetAlibabaMjMosFundModifybillbankaccountAPIRequest() *AlibabaMjMosFundModifybillbankaccountAPIRequest {
	return poolAlibabaMjMosFundModifybillbankaccountAPIRequest.Get().(*AlibabaMjMosFundModifybillbankaccountAPIRequest)
}

// ReleaseAlibabaMjMosFundModifybillbankaccountAPIRequest 将 AlibabaMjMosFundModifybillbankaccountAPIRequest 放入 sync.Pool
func ReleaseAlibabaMjMosFundModifybillbankaccountAPIRequest(v *AlibabaMjMosFundModifybillbankaccountAPIRequest) {
	v.Reset()
	poolAlibabaMjMosFundModifybillbankaccountAPIRequest.Put(v)
}
