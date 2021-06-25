package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量启动暂停推广词状态 APIRequest
alibaba.scbp.ad.keyword.status.batchupdate

批量启动暂停关键词推广状态
*/
type AlibabaScbpAdKeywordStatusBatchupdateRequest struct {
    model.Params

    // 系统自动生成
    keywordUpdateDtoList   []KeywordUpdateDto 

}

func NewAlibabaScbpAdKeywordStatusBatchupdateRequest() *AlibabaScbpAdKeywordStatusBatchupdateRequest{
    return &AlibabaScbpAdKeywordStatusBatchupdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdKeywordStatusBatchupdateRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.status.batchupdate"
}

func (r AlibabaScbpAdKeywordStatusBatchupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdKeywordStatusBatchupdateRequest) SetKeywordUpdateDtoList(keywordUpdateDtoList []KeywordUpdateDto) error {
    r.keywordUpdateDtoList = keywordUpdateDtoList
    r.Set("keyword_update_dto_list", keywordUpdateDtoList)
    return nil
}

func (r AlibabaScbpAdKeywordStatusBatchupdateRequest) GetKeywordUpdateDtoList() []KeywordUpdateDto {
    return r.keywordUpdateDtoList
}

