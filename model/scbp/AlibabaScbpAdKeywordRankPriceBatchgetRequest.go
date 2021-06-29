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
    context   *ContextDto
    // keyword_request
    keywordRequest   *TopKeywordListDTO
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
func (r *AlibabaScbpAdKeywordRankPriceBatchgetRequest) SetContext(context *ContextDto) error {
    r.context = context
    r.Set("context", context)
    return nil
}

// Context Getter
func (r AlibabaScbpAdKeywordRankPriceBatchgetRequest) GetContext() *ContextDto {
    return r.context
}
// KeywordRequest Setter
// keyword_request
func (r *AlibabaScbpAdKeywordRankPriceBatchgetRequest) SetKeywordRequest(keywordRequest *TopKeywordListDTO) error {
    r.keywordRequest = keywordRequest
    r.Set("keyword_request", keywordRequest)
    return nil
}

// KeywordRequest Getter
func (r AlibabaScbpAdKeywordRankPriceBatchgetRequest) GetKeywordRequest() *TopKeywordListDTO {
    return r.keywordRequest
}
