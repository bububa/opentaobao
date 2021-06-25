package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词批量改价 APIRequest
alibaba.scbp.ad.keyword.price.batchupdate

关键词批量改价
*/
type AlibabaScbpAdKeywordPriceBatchupdateRequest struct {
    model.Params

    // 系统自动生成
    keywordUpdateDtoList   []KeywordUpdateDto 

}

func NewAlibabaScbpAdKeywordPriceBatchupdateRequest() *AlibabaScbpAdKeywordPriceBatchupdateRequest{
    return &AlibabaScbpAdKeywordPriceBatchupdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdKeywordPriceBatchupdateRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.price.batchupdate"
}

func (r AlibabaScbpAdKeywordPriceBatchupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdKeywordPriceBatchupdateRequest) SetKeywordUpdateDtoList(keywordUpdateDtoList []KeywordUpdateDto) error {
    r.keywordUpdateDtoList = keywordUpdateDtoList
    r.Set("keyword_update_dto_list", keywordUpdateDtoList)
    return nil
}

func (r AlibabaScbpAdKeywordPriceBatchupdateRequest) GetKeywordUpdateDtoList() []KeywordUpdateDto {
    return r.keywordUpdateDtoList
}

