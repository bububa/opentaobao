package nrt

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtPayMerchantFundstypeModifyAPIRequest 修改摊位分账类型 API请求
// tmall.nrt.pay.merchant.fundstype.modify
//
// 修改摊位分账类型
type TmallNrtPayMerchantFundstypeModifyAPIRequest struct {
	model.Params
	// 请求参数
	_modifyFundsTypeReqDto *ModifyFundsTypeReqDto
}

// NewTmallNrtPayMerchantFundstypeModifyRequest 初始化TmallNrtPayMerchantFundstypeModifyAPIRequest对象
func NewTmallNrtPayMerchantFundstypeModifyRequest() *TmallNrtPayMerchantFundstypeModifyAPIRequest {
	return &TmallNrtPayMerchantFundstypeModifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrtPayMerchantFundstypeModifyAPIRequest) Reset() {
	r._modifyFundsTypeReqDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtPayMerchantFundstypeModifyAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.pay.merchant.fundstype.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtPayMerchantFundstypeModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtPayMerchantFundstypeModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModifyFundsTypeReqDto is ModifyFundsTypeReqDto Setter
// 请求参数
func (r *TmallNrtPayMerchantFundstypeModifyAPIRequest) SetModifyFundsTypeReqDto(_modifyFundsTypeReqDto *ModifyFundsTypeReqDto) error {
	r._modifyFundsTypeReqDto = _modifyFundsTypeReqDto
	r.Set("modify_funds_type_req_dto", _modifyFundsTypeReqDto)
	return nil
}

// GetModifyFundsTypeReqDto ModifyFundsTypeReqDto Getter
func (r TmallNrtPayMerchantFundstypeModifyAPIRequest) GetModifyFundsTypeReqDto() *ModifyFundsTypeReqDto {
	return r._modifyFundsTypeReqDto
}

var poolTmallNrtPayMerchantFundstypeModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrtPayMerchantFundstypeModifyRequest()
	},
}

// GetTmallNrtPayMerchantFundstypeModifyRequest 从 sync.Pool 获取 TmallNrtPayMerchantFundstypeModifyAPIRequest
func GetTmallNrtPayMerchantFundstypeModifyAPIRequest() *TmallNrtPayMerchantFundstypeModifyAPIRequest {
	return poolTmallNrtPayMerchantFundstypeModifyAPIRequest.Get().(*TmallNrtPayMerchantFundstypeModifyAPIRequest)
}

// ReleaseTmallNrtPayMerchantFundstypeModifyAPIRequest 将 TmallNrtPayMerchantFundstypeModifyAPIRequest 放入 sync.Pool
func ReleaseTmallNrtPayMerchantFundstypeModifyAPIRequest(v *TmallNrtPayMerchantFundstypeModifyAPIRequest) {
	v.Reset()
	poolTmallNrtPayMerchantFundstypeModifyAPIRequest.Put(v)
}
