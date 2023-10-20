package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordpricebatchupdateAPIRequest 关键词批量改价 API请求
// alibaba.scbp.ad.keyword.price.batchupdate
//
// 关键词批量改价
type AlibabascbpadkeywordpricebatchupdateAPIRequest struct {
	model.Params
	// 系统自动生成
	_keywordUpdateDtoList []KeywordUpdateDto
}

// NewAlibabascbpadkeywordpricebatchupdateRequest 初始化AlibabascbpadkeywordpricebatchupdateAPIRequest对象
func NewAlibabascbpadkeywordpricebatchupdateRequest() *AlibabascbpadkeywordpricebatchupdateAPIRequest {
	return &AlibabascbpadkeywordpricebatchupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadkeywordpricebatchupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.price.batchupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadkeywordpricebatchupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadkeywordpricebatchupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeywordUpdateDtoList is KeywordUpdateDtoList Setter
// 系统自动生成
func (r *AlibabascbpadkeywordpricebatchupdateAPIRequest) SetKeywordUpdateDtoList(_keywordUpdateDtoList []KeywordUpdateDto) error {
	r._keywordUpdateDtoList = _keywordUpdateDtoList
	r.Set("keyword_update_dto_list", _keywordUpdateDtoList)
	return nil
}

// GetKeywordUpdateDtoList KeywordUpdateDtoList Getter
func (r AlibabascbpadkeywordpricebatchupdateAPIRequest) GetKeywordUpdateDtoList() []KeywordUpdateDto {
	return r._keywordUpdateDtoList
}
