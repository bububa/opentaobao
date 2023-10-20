package aedropshiper

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressLogisticsBuyerFreightCalculateAPIRequest 提供给买家使用的运费计算接口 API请求
// aliexpress.logistics.buyer.freight.calculate
//
// 提供给买家使用的运费计算接口
type AliexpressLogisticsBuyerFreightCalculateAPIRequest struct {
	model.Params
	// 运费计算请求参数
	_paramAeopFreightCalculateForBuyerDTO *AeopFreightCalculateForBuyerDto
}

// NewAliexpressLogisticsBuyerFreightCalculateRequest 初始化AliexpressLogisticsBuyerFreightCalculateAPIRequest对象
func NewAliexpressLogisticsBuyerFreightCalculateRequest() *AliexpressLogisticsBuyerFreightCalculateAPIRequest {
	return &AliexpressLogisticsBuyerFreightCalculateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressLogisticsBuyerFreightCalculateAPIRequest) Reset() {
	r._paramAeopFreightCalculateForBuyerDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressLogisticsBuyerFreightCalculateAPIRequest) GetApiMethodName() string {
	return "aliexpress.logistics.buyer.freight.calculate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressLogisticsBuyerFreightCalculateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressLogisticsBuyerFreightCalculateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAeopFreightCalculateForBuyerDTO is ParamAeopFreightCalculateForBuyerDTO Setter
// 运费计算请求参数
func (r *AliexpressLogisticsBuyerFreightCalculateAPIRequest) SetParamAeopFreightCalculateForBuyerDTO(_paramAeopFreightCalculateForBuyerDTO *AeopFreightCalculateForBuyerDto) error {
	r._paramAeopFreightCalculateForBuyerDTO = _paramAeopFreightCalculateForBuyerDTO
	r.Set("param_aeop_freight_calculate_for_buyer_d_t_o", _paramAeopFreightCalculateForBuyerDTO)
	return nil
}

// GetParamAeopFreightCalculateForBuyerDTO ParamAeopFreightCalculateForBuyerDTO Getter
func (r AliexpressLogisticsBuyerFreightCalculateAPIRequest) GetParamAeopFreightCalculateForBuyerDTO() *AeopFreightCalculateForBuyerDto {
	return r._paramAeopFreightCalculateForBuyerDTO
}

var poolAliexpressLogisticsBuyerFreightCalculateAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressLogisticsBuyerFreightCalculateRequest()
	},
}

// GetAliexpressLogisticsBuyerFreightCalculateRequest 从 sync.Pool 获取 AliexpressLogisticsBuyerFreightCalculateAPIRequest
func GetAliexpressLogisticsBuyerFreightCalculateAPIRequest() *AliexpressLogisticsBuyerFreightCalculateAPIRequest {
	return poolAliexpressLogisticsBuyerFreightCalculateAPIRequest.Get().(*AliexpressLogisticsBuyerFreightCalculateAPIRequest)
}

// ReleaseAliexpressLogisticsBuyerFreightCalculateAPIRequest 将 AliexpressLogisticsBuyerFreightCalculateAPIRequest 放入 sync.Pool
func ReleaseAliexpressLogisticsBuyerFreightCalculateAPIRequest(v *AliexpressLogisticsBuyerFreightCalculateAPIRequest) {
	v.Reset()
	poolAliexpressLogisticsBuyerFreightCalculateAPIRequest.Put(v)
}
