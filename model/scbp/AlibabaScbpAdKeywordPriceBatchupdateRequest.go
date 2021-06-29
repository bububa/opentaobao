package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词批量改价 API请求
alibaba.scbp.ad.keyword.price.batchupdate

关键词批量改价
*/
type AlibabaScbpAdKeywordPriceBatchupdateRequest struct {
    model.Params
    // 系统自动生成
    _keywordUpdateDtoList   []KeywordUpdateDto
}

// 初始化AlibabaScbpAdKeywordPriceBatchupdateRequest对象
func NewAlibabaScbpAdKeywordPriceBatchupdateRequest() *AlibabaScbpAdKeywordPriceBatchupdateRequest{
    return &AlibabaScbpAdKeywordPriceBatchupdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordPriceBatchupdateRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.price.batchupdate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordPriceBatchupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// KeywordUpdateDtoList Setter
// 系统自动生成
func (r *AlibabaScbpAdKeywordPriceBatchupdateRequest) SetKeywordUpdateDtoList(_keywordUpdateDtoList []KeywordUpdateDto) error {
    r._keywordUpdateDtoList = _keywordUpdateDtoList
    r.Set("keyword_update_dto_list", _keywordUpdateDtoList)
    return nil
}

// KeywordUpdateDtoList Getter
func (r AlibabaScbpAdKeywordPriceBatchupdateRequest) GetKeywordUpdateDtoList() []KeywordUpdateDto {
    return r._keywordUpdateDtoList
}
