package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车关键词前五名批量排价 API请求
alibaba.scbp.ad.keyword.rank.price.batchget

外贸直通车关键词前五名批量排价
*/
type AlibabaScbpAdKeywordRankPriceBatchgetRequest struct {
    model.Params
    // 上下文
    _context   *ContextDTO
    // keyword_request
    _keywordRequest   *TopKeywordListDTO
}

// 初始化AlibabaScbpAdKeywordRankPriceBatchgetRequest对象
func NewAlibabaScbpAdKeywordRankPriceBatchgetRequest() *AlibabaScbpAdKeywordRankPriceBatchgetRequest{
    return &AlibabaScbpAdKeywordRankPriceBatchgetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordRankPriceBatchgetRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.rank.price.batchget"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordRankPriceBatchgetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Context Setter
// 上下文
func (r *AlibabaScbpAdKeywordRankPriceBatchgetRequest) SetContext(_context *ContextDTO) error {
    r._context = _context
    r.Set("context", _context)
    return nil
}

// Context Getter
func (r AlibabaScbpAdKeywordRankPriceBatchgetRequest) GetContext() *ContextDTO {
    return r._context
}
// KeywordRequest Setter
// keyword_request
func (r *AlibabaScbpAdKeywordRankPriceBatchgetRequest) SetKeywordRequest(_keywordRequest *TopKeywordListDTO) error {
    r._keywordRequest = _keywordRequest
    r.Set("keyword_request", _keywordRequest)
    return nil
}

// KeywordRequest Getter
func (r AlibabaScbpAdKeywordRankPriceBatchgetRequest) GetKeywordRequest() *TopKeywordListDTO {
    return r._keywordRequest
}
