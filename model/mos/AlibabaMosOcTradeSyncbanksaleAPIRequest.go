package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosOcTradeSyncbanksaleAPIRequest 云闪付、银行卡销售数据回传接口 API请求
// alibaba.mos.oc.trade.syncbanksale
//
// 云闪付、银行卡销售数据回传
type AlibabaMosOcTradeSyncbanksaleAPIRequest struct {
	model.Params
	// pos云闪付、银行卡销售数据
	_posBankSaleInfoDto *PosBankSaleInfoDto
}

// NewAlibabaMosOcTradeSyncbanksaleRequest 初始化AlibabaMosOcTradeSyncbanksaleAPIRequest对象
func NewAlibabaMosOcTradeSyncbanksaleRequest() *AlibabaMosOcTradeSyncbanksaleAPIRequest {
	return &AlibabaMosOcTradeSyncbanksaleAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosOcTradeSyncbanksaleAPIRequest) Reset() {
	r._posBankSaleInfoDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosOcTradeSyncbanksaleAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.oc.trade.syncbanksale"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosOcTradeSyncbanksaleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosOcTradeSyncbanksaleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPosBankSaleInfoDto is PosBankSaleInfoDto Setter
// pos云闪付、银行卡销售数据
func (r *AlibabaMosOcTradeSyncbanksaleAPIRequest) SetPosBankSaleInfoDto(_posBankSaleInfoDto *PosBankSaleInfoDto) error {
	r._posBankSaleInfoDto = _posBankSaleInfoDto
	r.Set("pos_bank_sale_info_dto", _posBankSaleInfoDto)
	return nil
}

// GetPosBankSaleInfoDto PosBankSaleInfoDto Getter
func (r AlibabaMosOcTradeSyncbanksaleAPIRequest) GetPosBankSaleInfoDto() *PosBankSaleInfoDto {
	return r._posBankSaleInfoDto
}

var poolAlibabaMosOcTradeSyncbanksaleAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosOcTradeSyncbanksaleRequest()
	},
}

// GetAlibabaMosOcTradeSyncbanksaleRequest 从 sync.Pool 获取 AlibabaMosOcTradeSyncbanksaleAPIRequest
func GetAlibabaMosOcTradeSyncbanksaleAPIRequest() *AlibabaMosOcTradeSyncbanksaleAPIRequest {
	return poolAlibabaMosOcTradeSyncbanksaleAPIRequest.Get().(*AlibabaMosOcTradeSyncbanksaleAPIRequest)
}

// ReleaseAlibabaMosOcTradeSyncbanksaleAPIRequest 将 AlibabaMosOcTradeSyncbanksaleAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosOcTradeSyncbanksaleAPIRequest(v *AlibabaMosOcTradeSyncbanksaleAPIRequest) {
	v.Reset()
	poolAlibabaMosOcTradeSyncbanksaleAPIRequest.Put(v)
}
