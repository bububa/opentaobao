package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordStatusBatchupdateAPIRequest 批量启动暂停推广词状态 API请求
// alibaba.scbp.ad.keyword.status.batchupdate
//
// 批量启动暂停关键词推广状态
type AlibabaScbpAdKeywordStatusBatchupdateAPIRequest struct {
	model.Params
	// 系统自动生成
	_keywordUpdateDtoList []KeywordUpdateDto
}

// NewAlibabaScbpAdKeywordStatusBatchupdateRequest 初始化AlibabaScbpAdKeywordStatusBatchupdateAPIRequest对象
func NewAlibabaScbpAdKeywordStatusBatchupdateRequest() *AlibabaScbpAdKeywordStatusBatchupdateAPIRequest {
	return &AlibabaScbpAdKeywordStatusBatchupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordStatusBatchupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.status.batchupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordStatusBatchupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdKeywordStatusBatchupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeywordUpdateDtoList is KeywordUpdateDtoList Setter
// 系统自动生成
func (r *AlibabaScbpAdKeywordStatusBatchupdateAPIRequest) SetKeywordUpdateDtoList(_keywordUpdateDtoList []KeywordUpdateDto) error {
	r._keywordUpdateDtoList = _keywordUpdateDtoList
	r.Set("keyword_update_dto_list", _keywordUpdateDtoList)
	return nil
}

// GetKeywordUpdateDtoList KeywordUpdateDtoList Getter
func (r AlibabaScbpAdKeywordStatusBatchupdateAPIRequest) GetKeywordUpdateDtoList() []KeywordUpdateDto {
	return r._keywordUpdateDtoList
}
