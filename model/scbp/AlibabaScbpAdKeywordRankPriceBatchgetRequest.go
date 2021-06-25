package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车关键词前五名批量排价 APIRequest
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

func NewAlibabaScbpAdKeywordRankPriceBatchgetRequest() *AlibabaScbpAdKeywordRankPriceBatchgetRequest{
    return &AlibabaScbpAdKeywordRankPriceBatchgetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdKeywordRankPriceBatchgetRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.rank.price.batchget"
}

func (r AlibabaScbpAdKeywordRankPriceBatchgetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdKeywordRankPriceBatchgetRequest) SetContext(context *ContextDto) error {
    r.context = context
    r.Set("context", context)
    return nil
}

func (r AlibabaScbpAdKeywordRankPriceBatchgetRequest) GetContext() *ContextDto {
    return r.context
}

func (r *AlibabaScbpAdKeywordRankPriceBatchgetRequest) SetKeywordRequest(keywordRequest *TopKeywordListDTO) error {
    r.keywordRequest = keywordRequest
    r.Set("keyword_request", keywordRequest)
    return nil
}

func (r AlibabaScbpAdKeywordRankPriceBatchgetRequest) GetKeywordRequest() *TopKeywordListDTO {
    return r.keywordRequest
}

