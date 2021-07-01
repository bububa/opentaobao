package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车批量删除关键词 API请求
alibaba.scbp.ad.keyword.batchdelete

外贸直通车批量删除关键词
*/
type AlibabaScbpAdKeywordBatchdeleteAPIRequest struct {
    model.Params
    // 关键词Id列表
    _keywordIdList   []int64
}

// 初始化AlibabaScbpAdKeywordBatchdeleteAPIRequest对象
func NewAlibabaScbpAdKeywordBatchdeleteRequest() *AlibabaScbpAdKeywordBatchdeleteAPIRequest{
    return &AlibabaScbpAdKeywordBatchdeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordBatchdeleteAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.batchdelete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordBatchdeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// KeywordIdList Setter
// 关键词Id列表
func (r *AlibabaScbpAdKeywordBatchdeleteAPIRequest) SetKeywordIdList(_keywordIdList []int64) error {
    r._keywordIdList = _keywordIdList
    r.Set("keyword_id_list", _keywordIdList)
    return nil
}

// KeywordIdList Getter
func (r AlibabaScbpAdKeywordBatchdeleteAPIRequest) GetKeywordIdList() []int64 {
    return r._keywordIdList
}
