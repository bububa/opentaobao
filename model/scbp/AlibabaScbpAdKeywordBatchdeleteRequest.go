package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车批量删除关键词 APIRequest
alibaba.scbp.ad.keyword.batchdelete

外贸直通车批量删除关键词
*/
type AlibabaScbpAdKeywordBatchdeleteRequest struct {
    model.Params

    // 关键词Id列表
    keywordIdList   []int64 

}

func NewAlibabaScbpAdKeywordBatchdeleteRequest() *AlibabaScbpAdKeywordBatchdeleteRequest{
    return &AlibabaScbpAdKeywordBatchdeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdKeywordBatchdeleteRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.batchdelete"
}

func (r AlibabaScbpAdKeywordBatchdeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdKeywordBatchdeleteRequest) SetKeywordIdList(keywordIdList []int64) error {
    r.keywordIdList = keywordIdList
    r.Set("keyword_id_list", keywordIdList)
    return nil
}

func (r AlibabaScbpAdKeywordBatchdeleteRequest) GetKeywordIdList() []int64 {
    return r.keywordIdList
}

