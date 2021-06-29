package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改关键词所属分组 API请求
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

// 初始化AlibabaScbpAdKeywordTagUpdateRequest对象
func NewAlibabaScbpAdKeywordTagUpdateRequest() *AlibabaScbpAdKeywordTagUpdateRequest{
    return &AlibabaScbpAdKeywordTagUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordTagUpdateRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.tag.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordTagUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// KeywordIdList Setter
// 关键词ID列表
func (r *AlibabaScbpAdKeywordTagUpdateRequest) SetKeywordIdList(keywordIdList []int64) error {
    r.keywordIdList = keywordIdList
    r.Set("keyword_id_list", keywordIdList)
    return nil
}

// KeywordIdList Getter
func (r AlibabaScbpAdKeywordTagUpdateRequest) GetKeywordIdList() []int64 {
    return r.keywordIdList
}
// TagIdList Setter
// 关键词分组ID,不传表示取消关键词的分组
func (r *AlibabaScbpAdKeywordTagUpdateRequest) SetTagIdList(tagIdList []int64) error {
    r.tagIdList = tagIdList
    r.Set("tag_id_list", tagIdList)
    return nil
}

// TagIdList Getter
func (r AlibabaScbpAdKeywordTagUpdateRequest) GetTagIdList() []int64 {
    return r.tagIdList
}
