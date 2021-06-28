package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改关键词所属分组 APIRequest
alibaba.scbp.ad.keyword.tag.update

修改关键词所属分组
*/
type AlibabaScbpAdKeywordTagUpdateRequest struct {
    model.Params

    // 关键词ID列表
    keywordIdList   []int64 

    // 关键词分组ID,不传表示取消关键词的分组
    tagIdList   []int64 

}

func NewAlibabaScbpAdKeywordTagUpdateRequest() *AlibabaScbpAdKeywordTagUpdateRequest{
    return &AlibabaScbpAdKeywordTagUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdKeywordTagUpdateRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.tag.update"
}

func (r AlibabaScbpAdKeywordTagUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdKeywordTagUpdateRequest) SetKeywordIdList(keywordIdList []int64) error {
    r.keywordIdList = keywordIdList
    r.Set("keyword_id_list", keywordIdList)
    return nil
}

func (r AlibabaScbpAdKeywordTagUpdateRequest) GetKeywordIdList() []int64 {
    return r.keywordIdList
}

func (r *AlibabaScbpAdKeywordTagUpdateRequest) SetTagIdList(tagIdList []int64) error {
    r.tagIdList = tagIdList
    r.Set("tag_id_list", tagIdList)
    return nil
}

func (r AlibabaScbpAdKeywordTagUpdateRequest) GetTagIdList() []int64 {
    return r.tagIdList
}

