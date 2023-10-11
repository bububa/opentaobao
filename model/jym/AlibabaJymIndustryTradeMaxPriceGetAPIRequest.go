package jym

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymIndustryTradeMaxPriceGetAPIRequest 获取交易猫单个游戏渠道帐号交易成功最高价 API请求
// alibaba.jym.industry.trade.max.price.get
//
// 获取交易猫单个游戏渠道帐号交易成功最高价
type AlibabaJymIndustryTradeMaxPriceGetAPIRequest struct {
	model.Params
	// 请求
	_jymMaxPriceRequest *JymMaxPriceRequestDto
}

// NewAlibabaJymIndustryTradeMaxPriceGetRequest 初始化AlibabaJymIndustryTradeMaxPriceGetAPIRequest对象
func NewAlibabaJymIndustryTradeMaxPriceGetRequest() *AlibabaJymIndustryTradeMaxPriceGetAPIRequest {
	return &AlibabaJymIndustryTradeMaxPriceGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymIndustryTradeMaxPriceGetAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.industry.trade.max.price.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymIndustryTradeMaxPriceGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJymIndustryTradeMaxPriceGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetJymMaxPriceRequest is JymMaxPriceRequest Setter
// 请求
func (r *AlibabaJymIndustryTradeMaxPriceGetAPIRequest) SetJymMaxPriceRequest(_jymMaxPriceRequest *JymMaxPriceRequestDto) error {
	r._jymMaxPriceRequest = _jymMaxPriceRequest
	r.Set("jym_max_price_request", _jymMaxPriceRequest)
	return nil
}

// GetJymMaxPriceRequest JymMaxPriceRequest Getter
func (r AlibabaJymIndustryTradeMaxPriceGetAPIRequest) GetJymMaxPriceRequest() *JymMaxPriceRequestDto {
	return r._jymMaxPriceRequest
}
