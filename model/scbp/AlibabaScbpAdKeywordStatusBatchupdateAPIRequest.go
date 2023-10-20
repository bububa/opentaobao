package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordstatusbatchupdateAPIRequest 批量启动暂停推广词状态 API请求
// alibaba.scbp.ad.keyword.status.batchupdate
//
// 批量启动暂停关键词推广状态
type AlibabascbpadkeywordstatusbatchupdateAPIRequest struct {
	model.Params
	// 系统自动生成
	_keywordUpdateDtoList []KeywordUpdateDto
}

// NewAlibabascbpadkeywordstatusbatchupdateRequest 初始化AlibabascbpadkeywordstatusbatchupdateAPIRequest对象
func NewAlibabascbpadkeywordstatusbatchupdateRequest() *AlibabascbpadkeywordstatusbatchupdateAPIRequest {
	return &AlibabascbpadkeywordstatusbatchupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadkeywordstatusbatchupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.status.batchupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadkeywordstatusbatchupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadkeywordstatusbatchupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeywordUpdateDtoList is KeywordUpdateDtoList Setter
// 系统自动生成
func (r *AlibabascbpadkeywordstatusbatchupdateAPIRequest) SetKeywordUpdateDtoList(_keywordUpdateDtoList []KeywordUpdateDto) error {
	r._keywordUpdateDtoList = _keywordUpdateDtoList
	r.Set("keyword_update_dto_list", _keywordUpdateDtoList)
	return nil
}

// GetKeywordUpdateDtoList KeywordUpdateDtoList Getter
func (r AlibabascbpadkeywordstatusbatchupdateAPIRequest) GetKeywordUpdateDtoList() []KeywordUpdateDto {
	return r._keywordUpdateDtoList
}
