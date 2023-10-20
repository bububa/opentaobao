package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordPriceBatchupdateAPIRequest 关键词批量改价 API请求
// alibaba.scbp.ad.keyword.price.batchupdate
//
// 关键词批量改价
type AlibabaScbpAdKeywordPriceBatchupdateAPIRequest struct {
	model.Params
	// 系统自动生成
	_keywordUpdateDtoList []KeywordUpdateDto
}

// NewAlibabaScbpAdKeywordPriceBatchupdateRequest 初始化AlibabaScbpAdKeywordPriceBatchupdateAPIRequest对象
func NewAlibabaScbpAdKeywordPriceBatchupdateRequest() *AlibabaScbpAdKeywordPriceBatchupdateAPIRequest {
	return &AlibabaScbpAdKeywordPriceBatchupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordPriceBatchupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.price.batchupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordPriceBatchupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdKeywordPriceBatchupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeywordUpdateDtoList is KeywordUpdateDtoList Setter
// 系统自动生成
func (r *AlibabaScbpAdKeywordPriceBatchupdateAPIRequest) SetKeywordUpdateDtoList(_keywordUpdateDtoList []KeywordUpdateDto) error {
	r._keywordUpdateDtoList = _keywordUpdateDtoList
	r.Set("keyword_update_dto_list", _keywordUpdateDtoList)
	return nil
}

// GetKeywordUpdateDtoList KeywordUpdateDtoList Getter
func (r AlibabaScbpAdKeywordPriceBatchupdateAPIRequest) GetKeywordUpdateDtoList() []KeywordUpdateDto {
	return r._keywordUpdateDtoList
}
