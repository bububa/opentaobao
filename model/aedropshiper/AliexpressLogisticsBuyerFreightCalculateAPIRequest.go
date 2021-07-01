package aedropshiper

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressLogisticsBuyerFreightCalculateAPIRequest
提供给买家使用的运费计算接口 API请求
aliexpress.logistics.buyer.freight.calculate

提供给买家使用的运费计算接口 */
type AliexpressLogisticsBuyerFreightCalculateAPIRequest struct {
	model.Params
	// 运费计算请求参数
	_paramAeopFreightCalculateForBuyerDTO *AeopFreightCalculateForBuyerDto
}

// NewAliexpressLogisticsBuyerFreightCalculateRequest 初始化AliexpressLogisticsBuyerFreightCalculateAPIRequest对象
func NewAliexpressLogisticsBuyerFreightCalculateRequest() *AliexpressLogisticsBuyerFreightCalculateAPIRequest {
	return &AliexpressLogisticsBuyerFreightCalculateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressLogisticsBuyerFreightCalculateAPIRequest) GetApiMethodName() string {
	return "aliexpress.logistics.buyer.freight.calculate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressLogisticsBuyerFreightCalculateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamAeopFreightCalculateForBuyerDTO Setter
// 运费计算请求参数
func (r *AliexpressLogisticsBuyerFreightCalculateAPIRequest) SetParamAeopFreightCalculateForBuyerDTO(_paramAeopFreightCalculateForBuyerDTO *AeopFreightCalculateForBuyerDto) error {
	r._paramAeopFreightCalculateForBuyerDTO = _paramAeopFreightCalculateForBuyerDTO
	r.Set("param_aeop_freight_calculate_for_buyer_d_t_o", _paramAeopFreightCalculateForBuyerDTO)
	return nil
}

// Get ParamAeopFreightCalculateForBuyerDTO Getter
func (r AliexpressLogisticsBuyerFreightCalculateAPIRequest) GetParamAeopFreightCalculateForBuyerDTO() *AeopFreightCalculateForBuyerDto {
	return r._paramAeopFreightCalculateForBuyerDTO
}
