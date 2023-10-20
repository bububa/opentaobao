package scbp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdKeywordRankPriceGetAPIRequest) Reset() {
	r._keyword = ""
	r.Params.ToZero()
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

var poolAlibabaScbpAdKeywordRankPriceGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdKeywordRankPriceGetRequest()
	},
}

// GetAlibabaScbpAdKeywordRankPriceGetRequest 从 sync.Pool 获取 AlibabaScbpAdKeywordRankPriceGetAPIRequest
func GetAlibabaScbpAdKeywordRankPriceGetAPIRequest() *AlibabaScbpAdKeywordRankPriceGetAPIRequest {
	return poolAlibabaScbpAdKeywordRankPriceGetAPIRequest.Get().(*AlibabaScbpAdKeywordRankPriceGetAPIRequest)
}

// ReleaseAlibabaScbpAdKeywordRankPriceGetAPIRequest 将 AlibabaScbpAdKeywordRankPriceGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdKeywordRankPriceGetAPIRequest(v *AlibabaScbpAdKeywordRankPriceGetAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdKeywordRankPriceGetAPIRequest.Put(v)
}
