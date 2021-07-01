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
type AlibabaScbpAdKeywordTagUpdateAPIRequest struct {
    model.Params
    // 关键词ID列表
    _keywordIdList   []int64
    // 关键词分组ID,不传表示取消关键词的分组
    _tagIdList   []int64
}

// 初始化AlibabaScbpAdKeywordTagUpdateAPIRequest对象
func NewAlibabaScbpAdKeywordTagUpdateRequest() *AlibabaScbpAdKeywordTagUpdateAPIRequest{
    return &AlibabaScbpAdKeywordTagUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordTagUpdateAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.tag.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordTagUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// KeywordIdList Setter
// 关键词ID列表
func (r *AlibabaScbpAdKeywordTagUpdateAPIRequest) SetKeywordIdList(_keywordIdList []int64) error {
    r._keywordIdList = _keywordIdList
    r.Set("keyword_id_list", _keywordIdList)
    return nil
}

// KeywordIdList Getter
func (r AlibabaScbpAdKeywordTagUpdateAPIRequest) GetKeywordIdList() []int64 {
    return r._keywordIdList
}
// TagIdList Setter
// 关键词分组ID,不传表示取消关键词的分组
func (r *AlibabaScbpAdKeywordTagUpdateAPIRequest) SetTagIdList(_tagIdList []int64) error {
    r._tagIdList = _tagIdList
    r.Set("tag_id_list", _tagIdList)
    return nil
}

// TagIdList Getter
func (r AlibabaScbpAdKeywordTagUpdateAPIRequest) GetTagIdList() []int64 {
    return r._tagIdList
}
