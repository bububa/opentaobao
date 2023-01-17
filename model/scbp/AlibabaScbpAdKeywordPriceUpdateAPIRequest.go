package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordPriceUpdateAPIRequest 关键词改价 API请求
// alibaba.scbp.ad.keyword.price.update
//
// 关键词改价
type AlibabaScbpAdKeywordPriceUpdateAPIRequest struct {
	model.Params
	// 只能取ascci字符
	_adKeyword string
	// 关键词价格单位元，一位小数
	_priceStr string
}

// NewAlibabaScbpAdKeywordPriceUpdateRequest 初始化AlibabaScbpAdKeywordPriceUpdateAPIRequest对象
func NewAlibabaScbpAdKeywordPriceUpdateRequest() *AlibabaScbpAdKeywordPriceUpdateAPIRequest {
	return &AlibabaScbpAdKeywordPriceUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordPriceUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.price.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordPriceUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdKeywordPriceUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdKeyword is AdKeyword Setter
// 只能取ascci字符
func (r *AlibabaScbpAdKeywordPriceUpdateAPIRequest) SetAdKeyword(_adKeyword string) error {
	r._adKeyword = _adKeyword
	r.Set("ad_keyword", _adKeyword)
	return nil
}

// GetAdKeyword AdKeyword Getter
func (r AlibabaScbpAdKeywordPriceUpdateAPIRequest) GetAdKeyword() string {
	return r._adKeyword
}

// SetPriceStr is PriceStr Setter
// 关键词价格单位元，一位小数
func (r *AlibabaScbpAdKeywordPriceUpdateAPIRequest) SetPriceStr(_priceStr string) error {
	r._priceStr = _priceStr
	r.Set("price_str", _priceStr)
	return nil
}

// GetPriceStr PriceStr Getter
func (r AlibabaScbpAdKeywordPriceUpdateAPIRequest) GetPriceStr() string {
	return r._priceStr
}
