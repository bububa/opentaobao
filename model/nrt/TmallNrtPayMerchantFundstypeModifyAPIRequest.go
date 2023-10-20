package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtpaymerchantfundstypemodifyAPIRequest 修改摊位分账类型 API请求
// tmall.nrt.pay.merchant.fundstype.modify
//
// 修改摊位分账类型
type TmallnrtpaymerchantfundstypemodifyAPIRequest struct {
	model.Params
	// 请求参数
	_modifyFundsTypeReqDto *ModifyFundsTypeReqDto
}

// NewTmallnrtpaymerchantfundstypemodifyRequest 初始化TmallnrtpaymerchantfundstypemodifyAPIRequest对象
func NewTmallnrtpaymerchantfundstypemodifyRequest() *TmallnrtpaymerchantfundstypemodifyAPIRequest {
	return &TmallnrtpaymerchantfundstypemodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtpaymerchantfundstypemodifyAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.pay.merchant.fundstype.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtpaymerchantfundstypemodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtpaymerchantfundstypemodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModifyFundsTypeReqDto is ModifyFundsTypeReqDto Setter
// 请求参数
func (r *TmallnrtpaymerchantfundstypemodifyAPIRequest) SetModifyFundsTypeReqDto(_modifyFundsTypeReqDto *ModifyFundsTypeReqDto) error {
	r._modifyFundsTypeReqDto = _modifyFundsTypeReqDto
	r.Set("modify_funds_type_req_dto", _modifyFundsTypeReqDto)
	return nil
}

// GetModifyFundsTypeReqDto ModifyFundsTypeReqDto Getter
func (r TmallnrtpaymerchantfundstypemodifyAPIRequest) GetModifyFundsTypeReqDto() *ModifyFundsTypeReqDto {
	return r._modifyFundsTypeReqDto
}
