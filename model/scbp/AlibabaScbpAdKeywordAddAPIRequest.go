package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordAddAPIRequest 外贸直通车加词 API请求
// alibaba.scbp.ad.keyword.add
//
// 外贸直通车加词服务
type AlibabaScbpAdKeywordAddAPIRequest struct {
	model.Params
	// 加入的词
	_adKeyword string
	// 词的出价
	_priceStr string
	// 分组名
	_tagName string
}

// NewAlibabaScbpAdKeywordAddRequest 初始化AlibabaScbpAdKeywordAddAPIRequest对象
func NewAlibabaScbpAdKeywordAddRequest() *AlibabaScbpAdKeywordAddAPIRequest {
	return &AlibabaScbpAdKeywordAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordAddAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAdKeyword is AdKeyword Setter
// 加入的词
func (r *AlibabaScbpAdKeywordAddAPIRequest) SetAdKeyword(_adKeyword string) error {
	r._adKeyword = _adKeyword
	r.Set("ad_keyword", _adKeyword)
	return nil
}

// GetAdKeyword AdKeyword Getter
func (r AlibabaScbpAdKeywordAddAPIRequest) GetAdKeyword() string {
	return r._adKeyword
}

// SetPriceStr is PriceStr Setter
// 词的出价
func (r *AlibabaScbpAdKeywordAddAPIRequest) SetPriceStr(_priceStr string) error {
	r._priceStr = _priceStr
	r.Set("price_str", _priceStr)
	return nil
}

// GetPriceStr PriceStr Getter
func (r AlibabaScbpAdKeywordAddAPIRequest) GetPriceStr() string {
	return r._priceStr
}

// SetTagName is TagName Setter
// 分组名
func (r *AlibabaScbpAdKeywordAddAPIRequest) SetTagName(_tagName string) error {
	r._tagName = _tagName
	r.Set("tag_name", _tagName)
	return nil
}

// GetTagName TagName Getter
func (r AlibabaScbpAdKeywordAddAPIRequest) GetTagName() string {
	return r._tagName
}
