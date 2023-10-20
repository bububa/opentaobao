package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosoctradesyncbanksaleAPIRequest 云闪付、银行卡销售数据回传接口 API请求
// alibaba.mos.oc.trade.syncbanksale
//
// 云闪付、银行卡销售数据回传
type AlibabamosoctradesyncbanksaleAPIRequest struct {
	model.Params
	// pos云闪付、银行卡销售数据
	_posBankSaleInfoDto *PosBankSaleInfoDto
}

// NewAlibabamosoctradesyncbanksaleRequest 初始化AlibabamosoctradesyncbanksaleAPIRequest对象
func NewAlibabamosoctradesyncbanksaleRequest() *AlibabamosoctradesyncbanksaleAPIRequest {
	return &AlibabamosoctradesyncbanksaleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosoctradesyncbanksaleAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.oc.trade.syncbanksale"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosoctradesyncbanksaleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosoctradesyncbanksaleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPosBankSaleInfoDto is PosBankSaleInfoDto Setter
// pos云闪付、银行卡销售数据
func (r *AlibabamosoctradesyncbanksaleAPIRequest) SetPosBankSaleInfoDto(_posBankSaleInfoDto *PosBankSaleInfoDto) error {
	r._posBankSaleInfoDto = _posBankSaleInfoDto
	r.Set("pos_bank_sale_info_dto", _posBankSaleInfoDto)
	return nil
}

// GetPosBankSaleInfoDto PosBankSaleInfoDto Getter
func (r AlibabamosoctradesyncbanksaleAPIRequest) GetPosBankSaleInfoDto() *PosBankSaleInfoDto {
	return r._posBankSaleInfoDto
}
