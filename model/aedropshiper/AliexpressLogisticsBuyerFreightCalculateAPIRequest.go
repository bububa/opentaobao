package aedropshiper

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresslogisticsbuyerfreightcalculateAPIRequest 提供给买家使用的运费计算接口 API请求
// aliexpress.logistics.buyer.freight.calculate
//
// 提供给买家使用的运费计算接口
type AliexpresslogisticsbuyerfreightcalculateAPIRequest struct {
	model.Params
	// 运费计算请求参数
	_paramAeopFreightCalculateForBuyerDTO *AeopFreightCalculateForBuyerDto
}

// NewAliexpresslogisticsbuyerfreightcalculateRequest 初始化AliexpresslogisticsbuyerfreightcalculateAPIRequest对象
func NewAliexpresslogisticsbuyerfreightcalculateRequest() *AliexpresslogisticsbuyerfreightcalculateAPIRequest {
	return &AliexpresslogisticsbuyerfreightcalculateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresslogisticsbuyerfreightcalculateAPIRequest) GetApiMethodName() string {
	return "aliexpress.logistics.buyer.freight.calculate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresslogisticsbuyerfreightcalculateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresslogisticsbuyerfreightcalculateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAeopFreightCalculateForBuyerDTO is ParamAeopFreightCalculateForBuyerDTO Setter
// 运费计算请求参数
func (r *AliexpresslogisticsbuyerfreightcalculateAPIRequest) SetParamAeopFreightCalculateForBuyerDTO(_paramAeopFreightCalculateForBuyerDTO *AeopFreightCalculateForBuyerDto) error {
	r._paramAeopFreightCalculateForBuyerDTO = _paramAeopFreightCalculateForBuyerDTO
	r.Set("param_aeop_freight_calculate_for_buyer_d_t_o", _paramAeopFreightCalculateForBuyerDTO)
	return nil
}

// GetParamAeopFreightCalculateForBuyerDTO ParamAeopFreightCalculateForBuyerDTO Getter
func (r AliexpresslogisticsbuyerfreightcalculateAPIRequest) GetParamAeopFreightCalculateForBuyerDTO() *AeopFreightCalculateForBuyerDto {
	return r._paramAeopFreightCalculateForBuyerDTO
}
