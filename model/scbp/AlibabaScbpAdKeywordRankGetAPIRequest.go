package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordRankGetAPIRequest 获取外贸直通车关键词预估排名 API请求
// alibaba.scbp.ad.keyword.rank.get
//
// 获取外贸直通车关键词预估排名
type AlibabaScbpAdKeywordRankGetAPIRequest struct {
	model.Params
	// 查询预估排名的关键词
	_keyword string
}

// NewAlibabaScbpAdKeywordRankGetRequest 初始化AlibabaScbpAdKeywordRankGetAPIRequest对象
func NewAlibabaScbpAdKeywordRankGetRequest() *AlibabaScbpAdKeywordRankGetAPIRequest {
	return &AlibabaScbpAdKeywordRankGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordRankGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.rank.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordRankGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetKeyword is Keyword Setter
// 查询预估排名的关键词
func (r *AlibabaScbpAdKeywordRankGetAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlibabaScbpAdKeywordRankGetAPIRequest) GetKeyword() string {
	return r._keyword
}
