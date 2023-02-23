package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordRankPriceGetAPIRequest 外贸直通车关键词前五名排价 API请求
// alibaba.scbp.ad.keyword.rank.price.get
//
// 外贸直通车关键词前五名排价
type AlibabaScbpAdKeywordRankPriceGetAPIRequest struct {
	model.Params
	// 关键词
	_keyword string
}

// NewAlibabaScbpAdKeywordRankPriceGetRequest 初始化AlibabaScbpAdKeywordRankPriceGetAPIRequest对象
func NewAlibabaScbpAdKeywordRankPriceGetRequest() *AlibabaScbpAdKeywordRankPriceGetAPIRequest {
	return &AlibabaScbpAdKeywordRankPriceGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordRankPriceGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.rank.price.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordRankPriceGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdKeywordRankPriceGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 关键词
func (r *AlibabaScbpAdKeywordRankPriceGetAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlibabaScbpAdKeywordRankPriceGetAPIRequest) GetKeyword() string {
	return r._keyword
}
