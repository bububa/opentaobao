package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordaddAPIRequest 外贸直通车加词 API请求
// alibaba.scbp.ad.keyword.add
//
// 外贸直通车加词服务
type AlibabascbpadkeywordaddAPIRequest struct {
	model.Params
	// 加入的词
	_adKeyword string
	// 词的出价
	_priceStr string
	// 分组名
	_tagName string
}

// NewAlibabascbpadkeywordaddRequest 初始化AlibabascbpadkeywordaddAPIRequest对象
func NewAlibabascbpadkeywordaddRequest() *AlibabascbpadkeywordaddAPIRequest {
	return &AlibabascbpadkeywordaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadkeywordaddAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadkeywordaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadkeywordaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdKeyword is AdKeyword Setter
// 加入的词
func (r *AlibabascbpadkeywordaddAPIRequest) SetAdKeyword(_adKeyword string) error {
	r._adKeyword = _adKeyword
	r.Set("ad_keyword", _adKeyword)
	return nil
}

// GetAdKeyword AdKeyword Getter
func (r AlibabascbpadkeywordaddAPIRequest) GetAdKeyword() string {
	return r._adKeyword
}

// SetPriceStr is PriceStr Setter
// 词的出价
func (r *AlibabascbpadkeywordaddAPIRequest) SetPriceStr(_priceStr string) error {
	r._priceStr = _priceStr
	r.Set("price_str", _priceStr)
	return nil
}

// GetPriceStr PriceStr Getter
func (r AlibabascbpadkeywordaddAPIRequest) GetPriceStr() string {
	return r._priceStr
}

// SetTagName is TagName Setter
// 分组名
func (r *AlibabascbpadkeywordaddAPIRequest) SetTagName(_tagName string) error {
	r._tagName = _tagName
	r.Set("tag_name", _tagName)
	return nil
}

// GetTagName TagName Getter
func (r AlibabascbpadkeywordaddAPIRequest) GetTagName() string {
	return r._tagName
}
