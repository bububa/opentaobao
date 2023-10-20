package scbp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdKeywordRankGetAPIRequest) Reset() {
	r._keyword = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordRankGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.rank.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordRankGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdKeywordRankGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaScbpAdKeywordRankGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdKeywordRankGetRequest()
	},
}

// GetAlibabaScbpAdKeywordRankGetRequest 从 sync.Pool 获取 AlibabaScbpAdKeywordRankGetAPIRequest
func GetAlibabaScbpAdKeywordRankGetAPIRequest() *AlibabaScbpAdKeywordRankGetAPIRequest {
	return poolAlibabaScbpAdKeywordRankGetAPIRequest.Get().(*AlibabaScbpAdKeywordRankGetAPIRequest)
}

// ReleaseAlibabaScbpAdKeywordRankGetAPIRequest 将 AlibabaScbpAdKeywordRankGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdKeywordRankGetAPIRequest(v *AlibabaScbpAdKeywordRankGetAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdKeywordRankGetAPIRequest.Put(v)
}
