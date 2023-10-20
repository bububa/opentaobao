package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjmosfundmodifybillbankaccountAPIRequest 修改付款单的银行账户信息 API请求
// alibaba.mj.mos.fund.modifybillbankaccount
//
// 修改付款单的银行账户信息
type AlibabamjmosfundmodifybillbankaccountAPIRequest struct {
	model.Params
	// 修改入参
	_modifyDto *ModifyBillDto
}

// NewAlibabamjmosfundmodifybillbankaccountRequest 初始化AlibabamjmosfundmodifybillbankaccountAPIRequest对象
func NewAlibabamjmosfundmodifybillbankaccountRequest() *AlibabamjmosfundmodifybillbankaccountAPIRequest {
	return &AlibabamjmosfundmodifybillbankaccountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamjmosfundmodifybillbankaccountAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.mos.fund.modifybillbankaccount"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamjmosfundmodifybillbankaccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamjmosfundmodifybillbankaccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModifyDto is ModifyDto Setter
// 修改入参
func (r *AlibabamjmosfundmodifybillbankaccountAPIRequest) SetModifyDto(_modifyDto *ModifyBillDto) error {
	r._modifyDto = _modifyDto
	r.Set("modify_dto", _modifyDto)
	return nil
}

// GetModifyDto ModifyDto Getter
func (r AlibabamjmosfundmodifybillbankaccountAPIRequest) GetModifyDto() *ModifyBillDto {
	return r._modifyDto
}
